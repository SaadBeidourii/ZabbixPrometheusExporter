package main

import (
	"log"
	"net/http"
	"time"

	zabbix "github.com/SaadBeidourii/ZabbixPrometheusExporter/internal/zabbix"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		for {
			zabbix.UpdateMetrics()
			time.Sleep(30 * time.Second) // Adjust the scrape interval as needed
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9100", nil)) // Expose metrics on port 9100
}
