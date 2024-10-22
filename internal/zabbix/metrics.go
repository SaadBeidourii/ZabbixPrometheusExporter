package zabbix

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Declare the database connection as a global variable
	db *sql.DB

	// Prometheus metric
	zabbixMetrics = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "zabbix_metric_name",
			Help: "Description of your Zabbix metric.",
		},
		[]string{"label1", "label2"},
	)
)

func init() {
	var err error
	// Initialize the database connection
	db, err = connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Register Prometheus metrics
	prometheus.MustRegister(zabbixMetrics)
}

func UpdateMetrics() {
	// Query the list of tables from the database
	tables, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("Failed to query tables: %v", err)
	}
	defer tables.Close()

	// Loop through all tables and log their names
	for tables.Next() {
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			log.Println("Error scanning table name:", err)
			continue
		}

		// Log the table name
		log.Printf("Table: %s", tableName)
	}

	// Check for errors from iterating over rows
	if err := tables.Err(); err != nil {
		log.Println("Error during iteration over tables:", err)
	}
}
