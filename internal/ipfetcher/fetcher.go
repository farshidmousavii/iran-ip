package ipfetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

func GetASN() ([]string, []string, []string, error) {
	log.Println("fetching ASN list from RIPE...")
	start := time.Now()

	resp, err := client.Get("https://stat.ripe.net/data/country-resource-list/data.json?resource=IR&v4_format=prefix&v6_format=prefix")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("http get failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, nil, fmt.Errorf("RIPE API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("can not read body: %w", err)
	}

	var asn ASNResponse
	if err := json.Unmarshal(body, &asn); err != nil {
		return nil, nil, nil, fmt.Errorf("can not conver data to json: %w", err)
	}

	log.Printf("found %d ASNs, %d IPv4 ranges, %d IPv6 ranges in %.2fs", len(asn.Data.Resources.ASN), len(asn.Data.Resources.IPV4), len(asn.Data.Resources.IPV6), time.Since(start).Seconds())
	return asn.Data.Resources.ASN, asn.Data.Resources.IPV4, asn.Data.Resources.IPV6, nil
}

func GetPrefixWorker(id int, jobs <-chan ASNJob, results chan<- PrefixResult) {
	const maxRetries = 5

	for job := range jobs {
		url := fmt.Sprintf("https://stat.ripe.net/data/announced-prefixes/data.json?resource=%s", job.ASN)

		var pre PrefixResponse
		var lastErr error
		ok := false

		for attempt := 0; attempt <= maxRetries; attempt++ {
			if attempt > 0 {
				backoff := time.Duration(1<<(attempt-1)) * time.Second
				log.Printf("worker %d: retrying %s in %s (attempt %d/%d)", id, job.ASN, backoff, attempt, maxRetries)
				time.Sleep(backoff)
			}

			resp, err := client.Get(url)
			if err != nil {
				lastErr = err
				continue
			}

			if resp.StatusCode != http.StatusOK {
				lastErr = fmt.Errorf("RIPE API returned status %d", resp.StatusCode)
				resp.Body.Close()
				continue
			}

			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				lastErr = err
				continue
			}

			if err := json.Unmarshal(body, &pre); err != nil {
				lastErr = err
				continue
			}

			ok = true
			break
		}

		if !ok {
			results <- PrefixResult{ASN: job.ASN, Err: fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)}
			continue
		}

		for _, p := range pre.Data.Prefixes {
			results <- PrefixResult{ASN: job.ASN, Prefix: p}
		}
	}
}

func GetPrefixes(asnList []string, workerCount int) []Prefix {
	log.Printf("starting %d workers to fetch prefixes for %d ASNs...", workerCount, len(asnList))
	start := time.Now()

	jobs := make(chan ASNJob)
	results := make(chan PrefixResult)

	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			GetPrefixWorker(id, jobs, results)
		}(i)
	}

	go func() {
		for _, asn := range asnList {
			jobs <- ASNJob{ASN: asn}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var all []Prefix
	var failed int
	for r := range results {
		if r.Err != nil {
			log.Printf("worker: failed to fetch %s: %v", r.ASN, r.Err)
			failed++
			continue
		}
		all = append(all, r.Prefix)
	}

	log.Printf("collected %d prefixes from %d ASNs (%d failed) in %.2fs", len(all), len(asnList), failed, time.Since(start).Seconds())
	return all
}
