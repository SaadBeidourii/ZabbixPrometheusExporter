package zabbix

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	zabbixMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zabbix_metric_name",
			Help: "Description of your Zabbix metric.",
		},
		[]string{"label1", "label2"},
	)
)

func init() {
	prometheus.MustRegister(zabbixMetrics)
}

func UpdateMetrics() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Example query
	rows, err := db.Query("SELECT some_metric FROM items")
	if err != nil {
		log.Fatalf("Failed to query metrics: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var metricValue float64
		if err := rows.Scan(&metricValue); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		// Update the Prometheus metric
		zabbixMetrics.WithLabelValues("exampleLabel").Set(metricValue)
	}
}
