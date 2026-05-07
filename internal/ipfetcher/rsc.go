package ipfetcher

import "strings"

func RscBuilder(subnets []string, listName string) string {
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

func RscBuilderV6(subnets []string, listName string) string {
	return RscBuilder(subnets, listName)
}
