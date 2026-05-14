package formatter

import (
	"bytes"
	"fmt"
)

type OpenWRTFormatter struct{}

func (OpenWRTFormatter) Name() string { return "openwrt" }

func (f OpenWRTFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("# last fetch: %s\n", timestamp))
	buf.WriteString(`#!/bin/sh
# Iran IP lists for OpenWRT
# Copy this file to your router and run: sh /etc/iran-ip.sh

echo "Importing Iran IP lists..."

`)

	if len(v4) > 0 {
		buf.WriteString("# IPv4 subnets\n")
		buf.WriteString("ipset destroy iran-v4 2>/dev/null\n")
		buf.WriteString("ipset create iran-v4 hash:net family inet hashsize 1024 maxelem 65536\n")
		for _, s := range v4 {
			buf.WriteString(fmt.Sprintf("ipset add iran-v4 %s\n", s))
		}
		buf.WriteString("\n")
	}

	if len(v6) > 0 {
		buf.WriteString("# IPv6 subnets\n")
		buf.WriteString("ipset destroy iran-v6 2>/dev/null\n")
		buf.WriteString("ipset create iran-v6 hash:net family inet6 hashsize 1024 maxelem 65536\n")
		for _, s := range v6 {
			buf.WriteString(fmt.Sprintf("ipset add iran-v6 %s\n", s))
		}
		buf.WriteString("\n")
	}

	buf.WriteString(`echo "Done."
echo ""
echo "Example iptables usage:"
echo "  iptables -t mangle -A PREROUTING -m set --match-set iran-v4 dst -j MARK --set-mark 1"
echo "  ip6tables -t mangle -A PREROUTING -m set --match-set iran-v6 dst -j MARK --set-mark 1"
echo ""
echo "For nftables (fw4 / OpenWRT 22.03+), use nftables/iran.nft instead."
`)

	return []File{
		{Path: "openwrt/iran.sh", Content: buf.Bytes()},
	}, nil
}
