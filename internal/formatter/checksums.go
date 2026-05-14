package formatter

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func GenerateChecksums(distDir string) error {
	var files []string
	err := filepath.Walk(distDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || info.Name() == "checksums.txt" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return fmt.Errorf("walk dist dir: %w", err)
	}

	sort.Strings(files)

	var content string
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("read %s: %w", f, err)
		}
		sum := sha256.Sum256(data)
		rel, _ := filepath.Rel(distDir, f)
		content += fmt.Sprintf("%x  %s\n", sum, rel)
	}

	path := filepath.Join(distDir, "checksums.txt")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("write checksums.txt: %w", err)
	}
	log.Printf("checksums.txt created (%d files)", len(files))
	return nil
}
