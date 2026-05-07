package ipfetcher

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"sort"
)

func Merge(prefixes []Prefix, ips []string) []string {
	log.Printf("merging %d prefixes with %d IP ranges...", len(prefixes), len(ips))

	set := make(map[string]struct{})
	for _, p := range prefixes {
		set[p.Prefix] = struct{}{}
	}
	for _, ip := range ips {
		set[ip] = struct{}{}
	}

	var result []string
	for k := range set {
		result = append(result, k)
	}

	log.Printf("merged to %d unique entries", len(result))
	return result
}

func splitByProtocol(subnets []string) (v4, v6 []string) {
	for _, s := range subnets {
		_, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			continue
		}
		if ipnet.IP.To4() != nil {
			v4 = append(v4, s)
		} else {
			v6 = append(v6, s)
		}
	}
	return
}

func NormalizeSubnets(subnets []string) []string {
	var v4s, v6s []*net.IPNet
	for _, s := range subnets {
		_, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			continue
		}

		if ipnet.IP.To4() != nil {
			v4s = append(v4s, ipnet)
		} else {
			v6s = append(v6s, ipnet)
		}
	}

	sort.Slice(v4s, func(i, j int) bool {
		return bytes.Compare(v4s[i].IP, v4s[j].IP) < 0
	})
	sort.Slice(v6s, func(i, j int) bool {
		return bytes.Compare(v6s[i].IP, v6s[j].IP) < 0
	})

	result := make([]string, 0, len(v4s)+len(v6s))
	for _, n := range v4s {
		result = append(result, n.String())
	}
	for _, n := range v6s {
		result = append(result, n.String())
	}

	return result
}

func removeSubnets(nets []*net.IPNet, isV4 bool) []*net.IPNet {
	var result []*net.IPNet
	for _, candidate := range nets {
		if len(result) == 0 {
			result = append(result, candidate)
			continue
		}
		last := result[len(result)-1]

		if last.Contains(candidate.IP) {
			continue
		}
		result = append(result, candidate)
	}
	return result
}

func MergeCIDRs(subnets []string) []string {
	v4, v6 := splitByProtocol(subnets)
	v4Merged := mergeIPv4(v4)
	v6Merged := mergeIPv6(v6)

	return append(v4Merged, v6Merged...)
}

func MergeCIDRsV4(subnets []string) []string {
	return mergeIPv4(subnets)
}

func MergeCIDRsV6(subnets []string) []string {
	return mergeIPv6(subnets)
}

func mergeIPv4(subnets []string) []string {
	log.Printf("merging %d IPv4 subnets...", len(subnets))

	var nets []*net.IPNet
	for _, s := range subnets {
		_, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			continue
		}
		if ipnet.IP.To4() == nil {
			continue
		}
		nets = append(nets, ipnet)
	}

	sort.Slice(nets, func(i, j int) bool {
		if cmp := bytes.Compare(nets[i].IP, nets[j].IP); cmp != 0 {
			return cmp < 0
		}
		mi, _ := nets[i].Mask.Size()
		mj, _ := nets[j].Mask.Size()
		return mi < mj
	})

	before := len(nets)
	nets = removeSubnets(nets, true)

	adjMerged := 0
	changed := true
	for changed {
		changed = false
		var newNets []*net.IPNet

		for i := 0; i < len(nets); i++ {
			if i+1 < len(nets) && canMergeV4(nets[i], nets[i+1]) {
				pair := mergePairV4(nets[i], nets[i+1])
				newNets = append(newNets, pair)
				i++
				changed = true
				adjMerged++
			} else {
				newNets = append(newNets, nets[i])
			}
		}

		nets = newNets
	}

	log.Printf("IPv4: removed %d overlapping, merged %d adjacent subnets", before-len(nets), adjMerged)

	out := make([]string, 0, len(nets))
	for _, n := range nets {
		out = append(out, n.String())
	}

	return out
}

func mergeIPv6(subnets []string) []string {
	log.Printf("merging %d IPv6 subnets...", len(subnets))

	var nets []*net.IPNet
	for _, s := range subnets {
		_, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			continue
		}
		if ipnet.IP.To4() != nil {
			continue
		}
		nets = append(nets, ipnet)
	}

	sort.Slice(nets, func(i, j int) bool {
		if cmp := bytes.Compare(nets[i].IP, nets[j].IP); cmp != 0 {
			return cmp < 0
		}
		mi, _ := nets[i].Mask.Size()
		mj, _ := nets[j].Mask.Size()
		return mi < mj
	})

	before := len(nets)
	nets = removeSubnets(nets, false)

	adjMerged := 0
	changed := true
	for changed {
		changed = false
		var newNets []*net.IPNet

		for i := 0; i < len(nets); i++ {
			if i+1 < len(nets) && canMergeV6(nets[i], nets[i+1]) {
				pair := mergePairV6(nets[i], nets[i+1])
				newNets = append(newNets, pair)
				i++
				changed = true
				adjMerged++
			} else {
				newNets = append(newNets, nets[i])
			}
		}

		nets = newNets
	}

	log.Printf("IPv6: removed %d overlapping, merged %d adjacent subnets", before-len(nets), adjMerged)

	out := make([]string, 0, len(nets))
	for _, n := range nets {
		out = append(out, n.String())
	}

	return out
}

func canMergeV4(a, b *net.IPNet) bool {
	ma, _ := a.Mask.Size()
	mb, _ := b.Mask.Size()

	if ma != mb {
		return false
	}

	size := 1 << (32 - ma)

	ai := binary.BigEndian.Uint32(a.IP.To4())
	bi := binary.BigEndian.Uint32(b.IP.To4())

	if bi-ai != uint32(size) {
		return false
	}

	superMask := net.CIDRMask(ma-1, 32)
	superIP := a.IP.Mask(superMask)

	return bytes.Equal(superIP, b.IP.Mask(superMask))
}

func mergePairV4(a, b *net.IPNet) *net.IPNet {
	maskSize, _ := a.Mask.Size()
	newMask := net.CIDRMask(maskSize-1, 32)

	ip := a.IP.Mask(newMask)

	return &net.IPNet{
		IP:   ip,
		Mask: newMask,
	}
}

func canMergeV6(a, b *net.IPNet) bool {
	ma, _ := a.Mask.Size()
	mb, _ := b.Mask.Size()

	if ma != mb {
		return false
	}

	aIP := a.IP.To16()
	bIP := b.IP.To16()
	if aIP == nil || bIP == nil {
		return false
	}

	carry := uint(0)
	diff := make([]byte, 16)
	for i := 15; i >= 0; i-- {
		d := int(bIP[i]) - int(aIP[i]) - int(carry)
		if d < 0 {
			d += 256
			carry = 1
		} else {
			carry = 0
		}
		diff[i] = byte(d)
	}

	subnetBits := 128 - ma
	sizeBytes := subnetBits / 8
	sizeRemainder := subnetBits % 8

	for i := 0; i < 16-sizeBytes-1; i++ {
		if diff[i] != 0 {
			return false
		}
	}

	targetByte := 16 - sizeBytes - 1
	expected := byte(1) << sizeRemainder
	if diff[targetByte] != expected {
		return false
	}

	superMask := net.CIDRMask(ma-1, 128)
	superA := a.IP.Mask(superMask)
	superB := b.IP.Mask(superMask)

	return bytes.Equal(superA, superB)
}

func mergePairV6(a, b *net.IPNet) *net.IPNet {
	maskSize, _ := a.Mask.Size()
	newMask := net.CIDRMask(maskSize-1, 128)

	ip := a.IP.Mask(newMask)

	return &net.IPNet{
		IP:   ip,
		Mask: newMask,
	}
}
