package formatter

import (
	"strings"
	"testing"
)

func TestNFTablesFormatter(t *testing.T) {
	f := NFTablesFormatter{}
	v4 := []string{"10.0.0.0/8", "192.168.0.0/16"}
	v6 := []string{"2001:db8::/32"}
	ts := "2024-01-01 12:00:00 UTC"

	files, err := f.Format(v4, v6, ts)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	// ipset file
	ipset := files[0]
	if ipset.Path != "firewall/iran.ipset" {
		t.Errorf("expected firewall/iran.ipset, got %s", ipset.Path)
	}
	ipsetContent := string(ipset.Content)
	if !strings.Contains(ipsetContent, "create iran-v4 hash:net family inet") {
		t.Error("missing create command for v4")
	}
	if !strings.Contains(ipsetContent, "create iran-v6 hash:net family inet6") {
		t.Error("missing create command for v6")
	}
	if !strings.Contains(ipsetContent, "add iran-v4 10.0.0.0/8") {
		t.Error("missing add for v4 subnet")
	}
	if !strings.Contains(ipsetContent, "add iran-v6 2001:db8::/32") {
		t.Error("missing add for v6 subnet")
	}

	// nft file
	nft := files[1]
	if nft.Path != "firewall/iran.nft" {
		t.Errorf("expected firewall/iran.nft, got %s", nft.Path)
	}
	nftContent := string(nft.Content)
	if !strings.Contains(nftContent, "table inet iran_ip") {
		t.Error("missing table declaration")
	}
	if !strings.Contains(nftContent, "type ipv4_addr") {
		t.Error("missing ipv4_addr type")
	}
	if !strings.Contains(nftContent, "type ipv6_addr") {
		t.Error("missing ipv6_addr type")
	}
	if !strings.Contains(nftContent, "flags interval") {
		t.Error("missing flags interval")
	}
}

func TestNFTablesFormatterV4Only(t *testing.T) {
	f := NFTablesFormatter{}
	v4 := []string{"10.0.0.0/8"}

	files, err := f.Format(v4, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}

	nftContent := string(files[1].Content)
	if strings.Contains(nftContent, "type ipv6_addr") {
		t.Error("should not contain ipv6_addr when no v6 subnets")
	}
}

func TestNFTablesFormatterEmpty(t *testing.T) {
	f := NFTablesFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}
}
