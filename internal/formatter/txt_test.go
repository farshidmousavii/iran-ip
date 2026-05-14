package formatter

import (
	"strings"
	"testing"
)

func TestTxtFormatter(t *testing.T) {
	f := TxtFormatter{}
	v4 := []string{"10.0.0.0/8", "192.168.0.0/16"}
	v6 := []string{"2001:db8::/32", "2a00::/29"}
	ts := "2024-01-01 12:00:00 UTC"

	files, err := f.Format(v4, v6, ts)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}

	v4f := files[0]
	if v4f.Path != "raw/ipv4.txt" {
		t.Errorf("expected raw/ipv4.txt, got %s", v4f.Path)
	}
	content := string(v4f.Content)
	if !strings.Contains(content, "10.0.0.0/8") {
		t.Error("missing 10.0.0.0/8 in output")
	}
	if !strings.Contains(content, "192.168.0.0/16") {
		t.Error("missing 192.168.0.0/16 in output")
	}
	if strings.Contains(content, "2001:db8::/32") {
		t.Error("IPv6 should not appear in IPv4 file")
	}
	if !strings.Contains(content, "# last fetch: "+ts) {
		t.Error("missing timestamp header")
	}

	v6f := files[1]
	if v6f.Path != "raw/ipv6.txt" {
		t.Errorf("expected raw/ipv6.txt, got %s", v6f.Path)
	}
	content6 := string(v6f.Content)
	if !strings.Contains(content6, "2001:db8::/32") {
		t.Error("missing 2001:db8::/32 in output")
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) != 3 { // header + 2 subnets
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
}

func TestTxtFormatterEmpty(t *testing.T) {
	f := TxtFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 files, got %d", len(files))
	}
	if len(files[0].Content) == 0 {
		t.Error("empty file should still have header")
	}
}
