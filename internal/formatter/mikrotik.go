package formatter

import (
	"fmt"
	"strings"
)

type MikrotikFormatter struct{}

func (MikrotikFormatter) Name() string { return "mikrotik" }

func (MikrotikFormatter) Format(v4, v6 []string, timestamp string) ([]File, error) {
	v4Content := fmt.Sprintf("# last fetch: %s\n", timestamp) + mikrotikRsc(v4, "IRAN")
	v6Content := fmt.Sprintf("# last fetch: %s\n", timestamp) + mikrotikRsc(v6, "IRANv6")

	return []File{
		{Path: "routeros/ipv4.rsc", Content: []byte(v4Content)},
		{Path: "routeros/ipv6.rsc", Content: []byte(v6Content)},
	}, nil
}

func mikrotikRsc(subnets []string, listName string) string {
	var result strings.Builder
	result.WriteString("/ip firewall address-list remove [/ip firewall address-list find list=" + listName + "]\n")
	result.WriteString("/ip firewall address-list\n")
	for _, m := range subnets {
		result.WriteString("add list=" + listName + " address=")
		result.WriteString(m)
		result.WriteString("\n")
	}
	return result.String()
}
