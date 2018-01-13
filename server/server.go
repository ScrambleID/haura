package server

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func New(options ...func(*Haura)) *Haura {
	s := &Haura{}

	for _, f := range options {
		f(s)
	}

	s.mux = http.NewServeMux()
	s.Logger = log.New(os.Stdout, "", 0)
	s.Stats = &Stats{StartTime: time.Now().Unix(), GoVersion: runtime.Version()}

	s.q = newQueue()

	s.mux.HandleFunc("/stats", s.stats)
	s.mux.HandleFunc("/enqueue", s.enqueue)
	s.mux.HandleFunc("/dequeue", s.dequeue)

	return s
}

func (s *Haura) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
