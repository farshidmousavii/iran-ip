package ipfetcher

import (
	"log"
	"path/filepath"
	"time"

	"github.com/farshidmousavii/iran-ip-ranges/internal/formatter"
)

func WriteFiles(subnets []string, dir string) error {
	v4, v6 := splitByProtocol(subnets)

	v4Merged := MergeCIDRsV4(v4)
	v6Merged := MergeCIDRsV6(v6)

	timestamp := time.Now().Format("2006-01-02 15:04:05 MST")
	tehran, err := time.LoadLocation("Asia/Tehran")
	if err == nil {
		timestamp = time.Now().In(tehran).Format("2006-01-02 15:04:05 MST")
	}

	_ = timestamp

	distDir := filepath.Join(dir, "dist")

	log.Println("generating dist/ files...")
	if err := formatter.RunAll(v4Merged, v6Merged, distDir); err != nil {
		return err
	}

	log.Println("generating checksums.txt...")
	if err := formatter.GenerateChecksums(distDir); err != nil {
		return err
	}

	return nil
}
