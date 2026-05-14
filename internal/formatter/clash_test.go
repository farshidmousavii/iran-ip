package formatter

import (
	"strings"
	"testing"
)

func TestClashFormatter(t *testing.T) {
	f := ClashFormatter{}
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
	if f0.Path != "clash/iran.yaml" {
		t.Errorf("expected clash/iran.yaml, got %s", f0.Path)
	}

	content := string(f0.Content)
	if !strings.Contains(content, "payload:") {
		t.Error("missing payload key")
	}
	if !strings.Contains(content, "'10.0.0.0/8'") {
		t.Error("missing 10.0.0.0/8")
	}
	if !strings.Contains(content, "'2001:db8::/32'") {
		t.Error("missing IPv6 subnet")
	}
	if !strings.Contains(content, "# last fetch: "+ts) {
		t.Error("missing timestamp")
	}

	lines := strings.Split(strings.TrimSpace(content), "\n")
	payloadCount := 0
	for _, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), "- ") {
			payloadCount++
		}
	}
	if payloadCount != 3 {
		t.Errorf("expected 3 payload entries, got %d", payloadCount)
	}
}

func TestClashFormatterEmpty(t *testing.T) {
	f := ClashFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	content := string(files[0].Content)
	if !strings.Contains(content, "payload:\n") {
		t.Error("empty file should still have payload key")
	}
}
