package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	h "github.com/vural/haura/server"
)

var hostAndPort = flag.String("listen", "0.0.0.0:8080", "ip and port to listen")

func init() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt)

	logger := log.New(os.Stdout, "", 0)

	s := h.New(func(s *h.Haura) { s.Logger = logger })
	s.Stats.Host = *hostAndPort

	h := &http.Server{Addr: *hostAndPort, Handler: s}

	go func() {
		logger.Printf("Listening on http://%s\n", *hostAndPort)

		if err := h.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	<-stop

	logger.Println("\nShutting down the haura...")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	h.Shutdown(ctx)

	logger.Println("Haura gracefully stopped")

}
