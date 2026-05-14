package formatter

import (
	"encoding/json"
	"testing"
)

func TestXrayFormatter(t *testing.T) {
	f := XrayFormatter{}
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
	if f0.Path != "xray/iran.json" {
		t.Errorf("expected xray/iran.json, got %s", f0.Path)
	}

	var result struct {
		Rules []struct {
			Type string   `json:"type"`
			IP   []string `json:"ip"`
		} `json:"rules"`
	}
	if err := json.Unmarshal(f0.Content, &result); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if len(result.Rules) != 1 {
		t.Fatalf("expected 1 rule, got %d", len(result.Rules))
	}
	if result.Rules[0].Type != "field" {
		t.Errorf("expected type field, got %s", result.Rules[0].Type)
	}
	if len(result.Rules[0].IP) != 3 {
		t.Fatalf("expected 3 IPs, got %d", len(result.Rules[0].IP))
	}
	if result.Rules[0].IP[0] != "10.0.0.0/8" {
		t.Errorf("expected 10.0.0.0/8, got %s", result.Rules[0].IP[0])
	}
}

func TestXrayFormatterEmpty(t *testing.T) {
	f := XrayFormatter{}
	files, err := f.Format(nil, nil, "ts")
	if err != nil {
		t.Fatal(err)
	}
	var result struct {
		Rules []interface{} `json:"rules"`
	}
	if err := json.Unmarshal(files[0].Content, &result); err != nil {
		t.Fatalf("invalid JSON for empty: %v", err)
	}
}
