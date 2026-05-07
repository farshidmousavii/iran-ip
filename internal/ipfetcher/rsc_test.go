package ipfetcher

import (
	"strings"
	"testing"
)

func TestRscBuilder(t *testing.T) {
	tests := []struct {
		name     string
		subnets  []string
		listName string
		contains []string
	}{
		{
			name:     "empty subnets",
			subnets:  nil,
			listName: "IRAN",
			contains: []string{
				"/ip firewall address-list remove",
				"/ip firewall address-list",
			},
		},
		{
			name:     "single subnet",
			subnets:  []string{"10.0.0.0/24"},
			listName: "IRAN",
			contains: []string{
				"add list=IRAN address=10.0.0.0/24",
			},
		},
		{
			name:     "multiple subnets",
			subnets:  []string{"10.0.0.0/24", "192.168.1.0/24"},
			listName: "TEST",
			contains: []string{
				"add list=TEST address=10.0.0.0/24",
				"add list=TEST address=192.168.1.0/24",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RscBuilder(tt.subnets, tt.listName)

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("missing %q in output", want)
				}
			}
		})
	}
}

func TestRscBuilderV6(t *testing.T) {
	tests := []struct {
		name     string
		subnets  []string
		listName string
		contains []string
	}{
		{
			name:     "empty subnets",
			subnets:  nil,
			listName: "IRANv6",
			contains: []string{
				"/ip firewall address-list remove",
				"/ip firewall address-list",
			},
		},
		{
			name:     "single subnet",
			subnets:  []string{"2001:db8::/32"},
			listName: "IRANv6",
			contains: []string{
				"add list=IRANv6 address=2001:db8::/32",
			},
		},
		{
			name:     "multiple subnets",
			subnets:  []string{"2001:db8::/32", "2a00::/29"},
			listName: "TEST6",
			contains: []string{
				"add list=TEST6 address=2001:db8::/32",
				"add list=TEST6 address=2a00::/29",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RscBuilderV6(tt.subnets, tt.listName)

			for _, want := range tt.contains {
				if !strings.Contains(result, want) {
					t.Errorf("missing %q in output", want)
				}
			}
		})
	}
}
