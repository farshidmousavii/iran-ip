package formatter

import (
	"encoding/json"
	"testing"
)

func TestRawJSONFormatter(t *testing.T) {
	f := RawJSONFormatter{}
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
	if f0.Path != "raw/iran.json" {
		t.Errorf("expected raw/iran.json, got %s", f0.Path)
	}

	var result []string
	if err := json.Unmarshal(f0.Content, &result); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if len(result) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(result))
	}
	if result[0] != "10.0.0.0/8" {
		t.Errorf("expected 10.0.0.0/8, got %s", result[0])
	}
}

func TestRawJSONFormatterEmpty(t *testing.T) {
	f := RawJSONFormatter{}
	files, err := f.Format(nil, nil, "")
	if err != nil {
		t.Fatal(err)
	}
	var result []string
	if err := json.Unmarshal(files[0].Content, &result); err != nil {
		t.Fatalf("invalid JSON for empty: %v", err)
	}
}
