package formatter

import (
	"encoding/json"
	"testing"
)

func TestSingboxFormatter(t *testing.T) {
	f := SingboxFormatter{}
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
	if f0.Path != "singbox/iran.json" {
		t.Errorf("expected singbox/iran.json, got %s", f0.Path)
	}

	var result struct {
		Version int `json:"version"`
		Rules   []struct {
			IPCIDR []string `json:"ip_cidr"`
		} `json:"rules"`
	}
	if err := json.Unmarshal(f0.Content, &result); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if result.Version != 3 {
		t.Errorf("expected version 3, got %d", result.Version)
	}
	if len(result.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(result.Rules))
	}
	if len(result.Rules[0].IPCIDR) != 3 {
		t.Fatalf("expected 3 IP CIDRs, got %d", len(result.Rules[0].IPCIDR))
	}
}

func TestSingboxFormatterEmpty(t *testing.T) {
	f := SingboxFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Version int `json:"version"`
	}
	if err := json.Unmarshal(files[0].Content, &result); err != nil {
		t.Fatalf("invalid JSON for empty: %v", err)
	}
}
