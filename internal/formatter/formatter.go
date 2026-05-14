package formatter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	Path    string
	Content []byte
}

type Formatter interface {
	Name() string
	Format(v4, v6 []string, timestamp string) ([]File, error)
}

func RunAll(v4, v6 []string, distDir string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05 MST")
	if loc, err := time.LoadLocation("Asia/Tehran"); err == nil {
		timestamp = time.Now().In(loc).Format("2006-01-02 15:04:05 MST")
	}

	formatters := []Formatter{
		TxtFormatter{},
		MikrotikFormatter{},
		ClashFormatter{},
		SingboxFormatter{},
		XrayFormatter{},
		NFTablesFormatter{},
		OpenWRTFormatter{},
		RawJSONFormatter{},
		RawYAMLFormatter{},
	}

	for _, f := range formatters {
		files, err := f.Format(v4, v6, timestamp)
		if err != nil {
			return fmt.Errorf("%s: %w", f.Name(), err)
		}
		for _, file := range files {
			path := filepath.Join(distDir, file.Path)
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				return fmt.Errorf("%s: mkdir %s: %w", f.Name(), filepath.Dir(file.Path), err)
			}
			if err := os.WriteFile(path, file.Content, 0644); err != nil {
				return fmt.Errorf("%s: write %s: %w", f.Name(), file.Path, err)
			}
			log.Printf("%s created (%d bytes)", file.Path, len(file.Content))
		}
	}

	return nil
}
