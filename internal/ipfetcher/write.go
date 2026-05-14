package ipfetcher

import (
	"log"
	"path/filepath"
	"time"

	"github.com/farshidmousavii/iran-ip/internal/formatter"
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

	log.Println("generating dist/ files...")
	if err := formatter.RunAll(v4Merged, v6Merged, filepath.Join(dir, "dist")); err != nil {
		return err
	}

	return nil
}
