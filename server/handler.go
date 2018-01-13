package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

func (s *Haura) enqueue(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method == "POST" {

		var params struct {
			Data string `json:"data"`
		}

		if err := json.NewDecoder(req.Body).Decode(&params); err != nil && params.Data != "" {
			http.Error(w, err.Error(), 400)
			return
		}

		s.q.Enqueue(params.Data)
		atomic.AddInt64(&s.Stats.TotalEnqueue, 1)
		atomic.AddInt64(&s.Stats.QueueCount, 1)

		fmt.Fprint(w, "ok")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func (s *Haura) dequeue(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method == "GET" {
		if val := s.q.Dequeue(); val != nil {
			atomic.AddInt64(&s.Stats.TotalDequeue, 1)
			atomic.AddInt64(&s.Stats.QueueCount, -1)
			fmt.Fprint(w, val)
			return
		}

		http.Error(w, "Queue is empty", 400)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func (s *Haura) stats(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.Stats)
}
