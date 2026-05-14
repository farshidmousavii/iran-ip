package formatter

import (
	"strings"
	"testing"
)

func TestOpenWRTFormatter(t *testing.T) {
	f := OpenWRTFormatter{}
	v4 := []string{"10.0.0.0/8", "192.168.0.0/16"}
	v6 := []string{"2001:db8::/32"}
	ts := "2024-01-01 12:00:00 UTC"

	files, err := f.Format(v4, v6, ts)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}

	f0 := files[0]
	if f0.Path != "openwrt/iran.sh" {
		t.Errorf("expected openwrt/iran.sh, got %s", f0.Path)
	}

	content := string(f0.Content)
	if !strings.Contains(content, "#!/bin/sh") {
		t.Error("missing shebang")
	}
	if !strings.Contains(content, "ipset create iran-v4 hash:net family inet") {
		t.Error("missing create command for v4")
	}
	if !strings.Contains(content, "ipset create iran-v6 hash:net family inet6") {
		t.Error("missing create command for v6")
	}
	if !strings.Contains(content, "ipset add iran-v4 10.0.0.0/8") {
		t.Error("missing add command for 10.0.0.0/8")
	}
	if !strings.Contains(content, "ipset add iran-v6 2001:db8::/32") {
		t.Error("missing add command for 2001:db8::/32")
	}
	if !strings.Contains(content, "#!/bin/sh") {
		t.Error("missing shebang")
	}
}

func TestOpenWRTFormatterEmpty(t *testing.T) {
	f := OpenWRTFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	content := string(files[0].Content)
	if !strings.Contains(content, "#!/bin/sh") {
		t.Error("missing shebang in empty output")
	}
}
