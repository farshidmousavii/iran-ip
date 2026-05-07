package web

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Server struct {
	dir         string
	refreshTick time.Duration
	mu          sync.RWMutex
	lastFetch   time.Time
	lastError   error
	srv         *http.Server
}

func New(dir string, refreshTick time.Duration, addr string) *Server {
	s := &Server{dir: dir, refreshTick: refreshTick}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/ipv4.txt", s.serveFile("ipv4.txt", "text/plain; charset=utf-8", false))
	mux.HandleFunc("/ipv6.txt", s.serveFile("ipv6.txt", "text/plain; charset=utf-8", false))
	mux.HandleFunc("/ipv4.rsc", s.serveFile("ipv4.rsc", "text/plain; charset=utf-8", true))
	mux.HandleFunc("/ipv6.rsc", s.serveFile("ipv6.rsc", "text/plain; charset=utf-8", true))

	s.srv = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	return s
}

func (s *Server) Start() error {
	go s.refreshLoop()

	log.Printf("web server started on %s", s.srv.Addr)
	log.Printf("endpoints:")
	log.Printf("  GET http://localhost%s/health     (health check)", s.srv.Addr)
	log.Printf("  GET http://localhost%s/ipv4.txt  (view in browser)", s.srv.Addr)
	log.Printf("  GET http://localhost%s/ipv6.txt  (view in browser)", s.srv.Addr)
	log.Printf("  GET http://localhost%s/ipv4.rsc  (download)", s.srv.Addr)
	log.Printf("  GET http://localhost%s/ipv6.rsc  (download)", s.srv.Addr)
	log.Printf("auto-refresh interval: %s", s.refreshTick)

	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) refreshLoop() {
	ticker := time.NewTicker(s.refreshTick)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("--- scheduled refresh ---")
		if err := fetchAndWrite(s.dir); err != nil {
			s.mu.Lock()
			s.lastFetch = time.Now()
			s.lastError = err
			s.mu.Unlock()
			log.Printf("refresh failed, serving cached files: %v", err)
		} else {
			s.mu.Lock()
			s.lastFetch = time.Now()
			s.lastError = nil
			s.mu.Unlock()
			log.Println("refresh complete, files updated")
		}
	}
}

func (s *Server) SetInitialFetch(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastFetch = time.Now()
	s.lastError = err
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	fetchTime := s.lastFetch
	err := s.lastError
	s.mu.RUnlock()

	if fetchTime.IsZero() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "initializing",
		})
		return
	}

	if err != nil {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusServiceUnavailable)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     "stale",
		"last_fetch": fetchTime.Format(time.RFC3339),
		"last_error": err.Error(),
	})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":     "ok",
		"last_fetch": fetchTime.Format(time.RFC3339),
	})
}

func (s *Server) serveFile(name, contentType string, download bool) http.HandlerFunc {
	path := filepath.Join(s.dir, name)
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("file request %s: not found", name)
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}

		log.Printf("serving %s to %s (modified: %s)", name, r.RemoteAddr, info.ModTime().Format(time.RFC1123))

		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=21600")
		if download {
			w.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")
		}
		w.Header().Set("Last-Modified", info.ModTime().Format(time.RFC1123))

		http.ServeFile(w, r, path)
	}
}
