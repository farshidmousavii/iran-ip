package ipfetcher

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		prefixes []Prefix
		ips      []string
		wantLen  int
	}{
		{
			name:     "empty inputs",
			prefixes: nil,
			ips:      nil,
			wantLen:  0,
		},
		{
			name: "deduplicates overlapping entries",
			prefixes: []Prefix{
				{Prefix: "10.0.0.0/24"},
				{Prefix: "192.168.1.0/24"},
			},
			ips: []string{
				"10.0.0.0/24",
				"172.16.0.0/16",
			},
			wantLen: 3,
		},
		{
			name: "handles IPv4 and IPv6 mixed",
			prefixes: []Prefix{
				{Prefix: "10.0.0.0/24"},
			},
			ips: []string{
				"2001:db8::/32",
				"192.168.1.0/24",
			},
			wantLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.prefixes, tt.ips)
			if len(result) != tt.wantLen {
				t.Errorf("got %d entries, want %d", len(result), tt.wantLen)
			}
		})
	}
}

func TestSplitByProtocol(t *testing.T) {
	tests := []struct {
		name    string
		subnets []string
		wantV4  int
		wantV6  int
	}{
		{
			name:    "empty",
			subnets: nil,
			wantV4:  0,
			wantV6:  0,
		},
		{
			name:    "IPv4 only",
			subnets: []string{"10.0.0.0/8", "192.168.0.0/16"},
			wantV4:  2,
			wantV6:  0,
		},
		{
			name:    "IPv6 only",
			subnets: []string{"2001:db8::/32", "2a00::/29"},
			wantV4:  0,
			wantV6:  2,
		},
		{
			name:    "mixed",
			subnets: []string{"10.0.0.0/8", "2001:db8::/32", "192.168.1.0/24", "2a00::/29"},
			wantV4:  2,
			wantV6:  2,
		},
		{
			name:    "invalid entries skipped",
			subnets: []string{"10.0.0.0/8", "invalid", "2001:db8::/32"},
			wantV4:  1,
			wantV6:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v4, v6 := splitByProtocol(tt.subnets)
			if len(v4) != tt.wantV4 {
				t.Errorf("v4: got %d, want %d", len(v4), tt.wantV4)
			}
			if len(v6) != tt.wantV6 {
				t.Errorf("v6: got %d, want %d", len(v6), tt.wantV6)
			}
		})
	}
}

func TestMergeCIDRsV4(t *testing.T) {
	tests := []struct {
		name    string
		subnets []string
		want    []string
	}{
		{
			name:    "empty",
			subnets: nil,
			want:    []string{},
		},
		{
			name:    "merges adjacent /24s into /23",
			subnets: []string{"10.0.0.0/24", "10.0.1.0/24"},
			want:    []string{"10.0.0.0/23"},
		},
		{
			name:    "removes contained subnets",
			subnets: []string{"10.0.0.0/8", "10.0.0.0/24"},
			want:    []string{"10.0.0.0/8"},
		},
		{
			name:    "no merge for non-adjacent",
			subnets: []string{"10.0.0.0/24", "10.0.2.0/24"},
			want:    []string{"10.0.0.0/24", "10.0.2.0/24"},
		},
		{
			name:    "ignores IPv6 entries",
			subnets: []string{"10.0.0.0/24", "2001:db8::/32"},
			want:    []string{"10.0.0.0/24"},
		},
		{
			name:    "multi-level merge",
			subnets: []string{"10.0.0.0/25", "10.0.0.128/25", "10.0.1.0/25", "10.0.1.128/25"},
			want:    []string{"10.0.0.0/23"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeCIDRsV4(tt.subnets)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}

func TestMergeCIDRsV6(t *testing.T) {
	tests := []struct {
		name    string
		subnets []string
		want    []string
	}{
		{
			name:    "empty",
			subnets: nil,
			want:    []string{},
		},
		{
			name:    "merges adjacent /49s into /48",
			subnets: []string{"2001:db8::/49", "2001:db8:0:8000::/49"},
			want:    []string{"2001:db8::/48"},
		},
		{
			name:    "removes contained subnets",
			subnets: []string{"2001:db8::/32", "2001:db8::/48"},
			want:    []string{"2001:db8::/32"},
		},
		{
			name:    "ignores IPv4 entries",
			subnets: []string{"2001:db8::/32", "10.0.0.0/8"},
			want:    []string{"2001:db8::/32"},
		},
		{
			name:    "no merge for non-adjacent",
			subnets: []string{"2001:db8::/32", "2002::/32"},
			want:    []string{"2001:db8::/32", "2002::/32"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeCIDRsV6(tt.subnets)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("got %v, want %v", result, tt.want)
			}
		})
	}
}

func TestMergeCIDRs(t *testing.T) {
	subnets := []string{
		"10.0.0.0/24", "10.0.1.0/24",
		"2001:db8::/49", "2001:db8:0:8000::/49",
	}

	result := MergeCIDRs(subnets)

	want := []string{"10.0.0.0/23", "2001:db8::/48"}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got %v, want %v", result, want)
	}
}

func TestNormalizeSubnets(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		wantLen int
	}{
		{
			name:    "sorts IPv4 correctly",
			input:   []string{"192.168.1.0/24", "10.0.0.0/8"},
			wantLen: 2,
		},
		{
			name:    "IPv4 before IPv6",
			input:   []string{"2001:db8::/32", "10.0.0.0/8"},
			wantLen: 2,
		},
		{
			name:    "skips invalid entries",
			input:   []string{"10.0.0.0/8", "not-a-cidr", "2001:db8::/32"},
			wantLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeSubnets(tt.input)
			if len(result) != tt.wantLen {
				t.Errorf("got %d entries, want %d", len(result), tt.wantLen)
			}
		})
	}
}
