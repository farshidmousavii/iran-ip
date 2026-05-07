package ipfetcher

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	log.Printf("creating ipv4.txt with %d subnets...", len(v4Merged))
	if err := writeTXT(filepath.Join(dir, "ipv4.txt"), v4Merged, timestamp, true); err != nil {
		return err
	}
	log.Printf("ipv4.txt created (%d entries)", len(v4Merged))

	log.Printf("creating ipv6.txt with %d subnets...", len(v6Merged))
	if err := writeTXT(filepath.Join(dir, "ipv6.txt"), v6Merged, timestamp, false); err != nil {
		return err
	}
	log.Printf("ipv6.txt created (%d entries)", len(v6Merged))

	log.Println("creating ipv4.rsc...")
	rscContent := "# last fetch: " + timestamp + "\n" + RscBuilder(v4Merged, "IRAN")
	if err := os.WriteFile(filepath.Join(dir, "ipv4.rsc"), []byte(rscContent), 0644); err != nil {
		return fmt.Errorf("cannot create ipv4.rsc: %w", err)
	}
	log.Printf("ipv4.rsc created")

	log.Println("creating ipv6.rsc...")
	rscV6Content := "# last fetch: " + timestamp + "\n" + RscBuilderV6(v6Merged, "IRANv6")
	if err := os.WriteFile(filepath.Join(dir, "ipv6.rsc"), []byte(rscV6Content), 0644); err != nil {
		return fmt.Errorf("cannot create ipv6.rsc: %w", err)
	}
	log.Printf("ipv6.rsc created")

	return nil
}

func writeTXT(path string, subnets []string, timestamp string, skipV6 bool) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot create %s: %w", filepath.Base(path), err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "# last fetch:", timestamp)
	for _, subnet := range subnets {
		if skipV6 && strings.Contains(subnet, ":") {
			continue
		}
		fmt.Fprintln(writer, subnet)
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush %s failed: %w", filepath.Base(path), err)
	}

	return nil
}
