package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func Serve(endpoint, port string) {
	port = ":" + port
	http.Handle(endpoint, promhttp.Handler())
	http.ListenAndServe(port, nil)
}
