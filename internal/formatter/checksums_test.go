package formatter

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateChecksums(t *testing.T) {
	dir := t.TempDir()

	// create test files
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("hello"), 0644)
	os.WriteFile(filepath.Join(dir, "b.txt"), []byte("world"), 0644)
	os.MkdirAll(filepath.Join(dir, "sub"), 0755)
	os.WriteFile(filepath.Join(dir, "sub", "c.txt"), []byte("test"), 0644)

	if err := GenerateChecksums(dir); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "checksums.txt"))
	if err != nil {
		t.Fatal(err)
	}

	content := string(data)
	if !strings.Contains(content, "a.txt") {
		t.Error("missing a.txt")
	}
	if !strings.Contains(content, "b.txt") {
		t.Error("missing b.txt")
	}
	if !strings.Contains(content, "sub/c.txt") {
		t.Error("missing sub/c.txt")
	}

	// verify SHA256 format
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}

	// each line should be: hex_hash  filename
	for _, line := range lines {
		if len(line) < 64 {
			t.Errorf("line too short for SHA256: %q", line)
		}
		if line[64] != ' ' || line[65] != ' ' {
			t.Errorf("expected double space after hash: %q", line)
		}
	}
}

func TestGenerateChecksumsSkipsSelf(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "checksums.txt"), []byte("old"), 0644)
	os.WriteFile(filepath.Join(dir, "a.txt"), []byte("data"), 0644)

	if err := GenerateChecksums(dir); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "checksums.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(data), "checksums.txt") {
		t.Error("checksums.txt should not list itself")
	}
}
