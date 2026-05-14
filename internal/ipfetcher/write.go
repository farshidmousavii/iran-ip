package ipfetcher

import (
	"log"
	"path/filepath"

	"github.com/farshidmousavii/iran-ip-ranges/internal/formatter"
)

func WriteFiles(subnets []string, dir string) error {
	v4, v6 := splitByProtocol(subnets)

	v4Merged := MergeCIDRsV4(v4)
	v6Merged := MergeCIDRsV6(v6)

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
