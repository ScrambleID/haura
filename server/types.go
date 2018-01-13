package server

import (
	"log"
	"net/http"
)

type Stats struct {
	GoVersion    string `json:"go_version"`
	Host         string `json:"host"`
	StartTime    int64  `json:"start_time"`
	TotalEnqueue int64  `json:"total_enqueue"`
	TotalDequeue int64  `json:"total_dequeue"`
	QueueCount   int64  `json:"queue_count"`
}

type Haura struct {
	Logger *log.Logger
	mux    *http.ServeMux
	q      *queue

	Stats *Stats
}
