// Author: Mitch Allen
// File: main.go

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/mitchallen/go-monorepo-demo/pkg/beta"
	"github.com/mitchallen/go-monorepo-demo/pkg/shared"
)

var logger *shared.Logger

func main() {
	logger = shared.NewLogger("web-server")

	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/flip", flipHandler)
	http.HandleFunc("/api/analyze", analyzeHandler)

	addr := fmt.Sprintf(":%d", *port)
	logger.Info(fmt.Sprintf("Starting web server on %s", addr))
	logger.Info("Available endpoints:")
	logger.Info("  GET  /        - Home")
	logger.Info("  GET  /health  - Health check")
	logger.Info("  GET  /api/flip?count=100 - Run coin flips")
	logger.Info("  GET  /api/analyze?count=100 - Analyze coin flips")

	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Error(fmt.Sprintf("Server failed: %v", err))
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Monorepo Demo API</title>
</head>
<body>
    <h1>Go Monorepo Demo - Web Server</h1>
    <h2>Available Endpoints:</h2>
    <ul>
        <li><a href="/health">/health</a> - Health check</li>
        <li><a href="/api/flip?count=100">/api/flip?count=100</a> - Run coin flips</li>
        <li><a href="/api/analyze?count=100">/api/analyze?count=100</a> - Analyze coin flips</li>
    </ul>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Health check requested")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func flipHandler(w http.ResponseWriter, r *http.Request) {
	count := getCountParam(r, 100)
	logger.Info(fmt.Sprintf("Flip requested with count=%d", count))

	analysis := beta.AnalyzeCoinFlips(count)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	count := getCountParam(r, 100)
	logger.Info(fmt.Sprintf("Analyze requested with count=%d", count))

	analysis := beta.AnalyzeCoinFlips(count)

	// Add percentage calculations
	result := map[string]interface{}{
		"total":          analysis["total"],
		"heads":          analysis["heads"],
		"tails":          analysis["tails"],
		"max":            analysis["max"],
		"min":            analysis["min"],
		"heads_percent":  float64(analysis["heads"]) / float64(analysis["total"]) * 100,
		"tails_percent":  float64(analysis["tails"]) / float64(analysis["total"]) * 100,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getCountParam(r *http.Request, defaultValue int) int {
	countStr := r.URL.Query().Get("count")
	if countStr == "" {
		return defaultValue
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		return defaultValue
	}

	return count
}
