package formatter

import (
	"strings"
	"testing"
)

func TestRawYAMLFormatter(t *testing.T) {
	f := RawYAMLFormatter{}
	v4 := []string{"10.0.0.0/8", "192.168.0.0/16"}
	v6 := []string{"2001:db8::/32"}

	files, err := f.Format(v4, v6, "")
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}

	f0 := files[0]
	if f0.Path != "raw/iran.yaml" {
		t.Errorf("expected raw/iran.yaml, got %s", f0.Path)
	}

	content := string(f0.Content)
	if !strings.Contains(content, `- "10.0.0.0/8"`) {
		t.Error("missing 10.0.0.0/8")
	}
	if !strings.Contains(content, `- "2001:db8::/32"`) {
		t.Error("missing IPv6")
	}
	if !strings.Contains(content, "# last fetch:") {
		t.Error("missing timestamp header")
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) != 4 {
		t.Errorf("expected 4 lines (header+3 entries), got %d", len(lines))
	}
}

func TestRawYAMLFormatterEmpty(t *testing.T) {
	f := RawYAMLFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	content := string(files[0].Content)
	if !strings.Contains(content, "# last fetch: ts") {
		t.Error("missing timestamp header")
	}
}
