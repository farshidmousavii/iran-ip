package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/farshidmousavii/iran-ip/internal/ipfetcher"
	"github.com/farshidmousavii/iran-ip/internal/web"
)

func main() {
	fetchOnly := flag.Bool("fetch-only", false, "fetch IPs, write files, and exit")
	addr := flag.String("addr", ":8080", "web server listen address")
	refresh := flag.Duration("refresh", 6*time.Hour, "auto-refresh interval for IP lists")
	flag.Parse()

	log.Println("========================================")
	log.Println(" iran-ip — Iran IPv4/IPv6 List Fetcher")
	log.Println("========================================")

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("working directory: %s", dir)

	if *fetchOnly {
		log.Println("mode: fetch-only")
		if err := fetchAndWrite(dir); err != nil {
			log.Fatalf("fetch failed: %v", err)
		}
		log.Println("done: all files created in dist/")
		return
	}

	log.Println("mode: fetch + serve")
	log.Println("attempting initial IP fetch...")

	srv := web.New(dir, *refresh, *addr)

	initErr := fetchAndWrite(dir)
	srv.SetInitialFetch(initErr)
	if initErr != nil {
		log.Printf("initial fetch failed: %v", initErr)
		log.Println("falling back to existing cached files")
	}

	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("received %s, shutting down...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}

	log.Println("server stopped")
}

func fetchAndWrite(dir string) error {
	asn, ipList, ipv6List, err := ipfetcher.GetASN()
	if err != nil {
		return err
	}

	prefixes := ipfetcher.GetPrefixes(asn, 50)
	subnets := ipfetcher.Merge(prefixes, append(ipList, ipv6List...))

	return ipfetcher.WriteFiles(subnets, dir)
}
