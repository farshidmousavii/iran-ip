package web

import (
	"log"

	"github.com/farshidmousavii/iran-ip-ranges/internal/ipfetcher"
)

// FetchAndWrite fetches Iran IP ranges from RIPE and writes all dist/ files.

func FetchAndWrite(dir string) error {
	log.Println("=== starting IP fetch cycle ===")

	asn, ipList, ipv6List, err := ipfetcher.GetASN()
	if err != nil {
		return err
	}

	prefixes := ipfetcher.GetPrefixes(asn, 50)
	subnets := ipfetcher.Merge(prefixes, append(ipList, ipv6List...))

	if err := ipfetcher.WriteFiles(subnets, dir); err != nil {
		return err
	}

	log.Printf("=== fetch cycle complete ===")
	return nil
}
