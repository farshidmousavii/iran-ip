package formatter

import (
	"bytes"
	"fmt"
)

type NFTablesFormatter struct{}

func (NFTablesFormatter) Name() string { return "nftables" }

func (f NFTablesFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	ipsetContent := f.buildIpset(v4, v6, timestamp)
	nftContent := f.buildNFT(v4, v6, timestamp)

	return []File{
		{Path: "nftables/iran.ipset", Content: ipsetContent},
		{Path: "nftables/iran.nft", Content: nftContent},
	}, nil
}

func (NFTablesFormatter) buildIpset(v4, v6 []string, timestamp string) []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("# last fetch: %s\n", timestamp))

	if len(v4) > 0 {
		buf.WriteString("# Iran IPv4 subnets\n")
		buf.WriteString("create iran-v4 hash:net family inet hashsize 1024 maxelem 65536\n")
		buf.WriteString("flush iran-v4\n")
		for _, s := range v4 {
			buf.WriteString(fmt.Sprintf("add iran-v4 %s\n", s))
		}
	}

	if len(v6) > 0 {
		buf.WriteString("# Iran IPv6 subnets\n")
		buf.WriteString("create iran-v6 hash:net family inet6 hashsize 1024 maxelem 65536\n")
		buf.WriteString("flush iran-v6\n")
		for _, s := range v6 {
			buf.WriteString(fmt.Sprintf("add iran-v6 %s\n", s))
		}
	}

	return buf.Bytes()
}

func (NFTablesFormatter) buildNFT(v4, v6 []string, timestamp string) []byte {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("# last fetch: %s\n", timestamp))
	buf.WriteString("# Iran IP subnets for nftables\n")
	buf.WriteString("# include this file:  include \"/etc/nftables/iran.nft\"\n")
	buf.WriteString("# or copy the sets into your existing table\n\n")
	buf.WriteString("table inet iran_ip {\n")

	if len(v4) > 0 {
		buf.WriteString("  set iran-v4 {\n")
		buf.WriteString("    type ipv4_addr\n")
		buf.WriteString("    flags interval\n")
		buf.WriteString("    auto-merge\n")
		buf.WriteString("    elements = {\n")
		for i, s := range v4 {
			comma := ","
			if i == len(v4)-1 && len(v6) == 0 {
				comma = ""
			}
			buf.WriteString(fmt.Sprintf("      %s%s\n", s, comma))
		}
		buf.WriteString("    }\n")
		buf.WriteString("  }\n")
	}

	if len(v6) > 0 {
		buf.WriteString("  set iran-v6 {\n")
		buf.WriteString("    type ipv6_addr\n")
		buf.WriteString("    flags interval\n")
		buf.WriteString("    auto-merge\n")
		buf.WriteString("    elements = {\n")
		for i, s := range v6 {
			comma := ","
			if i == len(v6)-1 {
				comma = ""
			}
			buf.WriteString(fmt.Sprintf("      %s%s\n", s, comma))
		}
		buf.WriteString("    }\n")
		buf.WriteString("  }\n")
	}

	buf.WriteString("}\n")
	return buf.Bytes()
}
