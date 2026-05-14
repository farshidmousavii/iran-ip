#!/bin/sh
# OpenWRT firewall configuration with Iran IP sets

# Copy iran.ipset to /etc/iran.ipset first
wget -O /etc/iran.ipset https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset
wget -O /etc/iran.nft https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft

# Import ipset (for iptables-based fw3)
ipset restore < /etc/iran.ipset

# Or for nftables-based fw4 (OpenWRT 22.03+):
nft -f /etc/iran.nft

# Add iptables rule to bypass VPN for Iran traffic
iptables -t mangle -A PREROUTING -m set --match-set iran-v4 dst -j MARK --set-mark 1
ip6tables -t mangle -A PREROUTING -m set --match-set iran-v6 dst -j MARK --set-mark 1

echo "Iran IP sets loaded successfully"
