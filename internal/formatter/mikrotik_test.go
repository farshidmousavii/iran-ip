package formatter

import (
	"strings"
	"testing"
)

func TestMikrotikFormatter(t *testing.T) {
	f := MikrotikFormatter{}
	v4 := []string{"10.0.0.0/24", "192.168.1.0/24"}
	v6 := []string{"2001:db8::/32"}
	ts := "2024-01-01 12:00:00 UTC"

	files, err := f.Format(v4, v6, ts)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	v4f := files[0]
	if v4f.Path != "mikrotik/ipv4.rsc" {
		t.Errorf("expected mikrotik/ipv4.rsc, got %s", v4f.Path)
	}
	content := string(v4f.Content)
	if !strings.Contains(content, "/ip firewall address-list remove") {
		t.Error("missing remove command")
	}
	if !strings.Contains(content, "/ip firewall address-list") {
		t.Error("missing address-list path")
	}
	if !strings.Contains(content, "add list=IRAN address=10.0.0.0/24") {
		t.Error("missing add command for subnet")
	}

	v6f := files[1]
	if v6f.Path != "mikrotik/ipv6.rsc" {
		t.Errorf("expected mikrotik/ipv6.rsc, got %s", v6f.Path)
	}
	content6 := string(v6f.Content)
	if !strings.Contains(content6, "add list=IRANv6 address=2001:db8::/32") {
		t.Error("missing add command for v6 subnet")
	}
}

func TestMikrotikFormatterEmpty(t *testing.T) {
	f := MikrotikFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}
}
