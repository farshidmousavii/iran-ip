#!/bin/sh
# Simple split tunneling with iptables + ipset

# Load Iran IP sets
wget -O /tmp/iran.ipset https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset
ipset restore < /tmp/iran.ipset

# Mark Iran traffic to bypass VPN
iptables -t mangle -A OUTPUT -m set --match-set iran-v4 dst -j MARK --set-mark 1
ip6tables -t mangle -A OUTPUT -m set --match-set iran-v6 dst -j MARK --set-mark 1

# Route marked traffic through main table
ip rule add fwmark 1 table main priority 100
