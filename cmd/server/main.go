package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/qweviluxx/GopherScanner.git/internal"
)

type ScanResponse struct {
	Hostname string `json:"hostname"`
	Ports    []int  `json:"ports"`
}

func validation(w http.ResponseWriter, h string, s, e int) bool {

	if h == "" || s < 0 || e < 0 || e < s || s > 65535 || e > 65535 {
		http.Error(w, "Bad params", http.StatusBadRequest)
		return false
	}

	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	scanner := internal.NewScanner("tcp")

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	params := r.URL.Query()

	hostname := params.Get("hostname")

	startPort, err := strconv.Atoi(params.Get("startport"))
	if err != nil {
		http.Error(w, "Parsing param error:", http.StatusBadRequest)
		return
	}

	endPort, err := strconv.Atoi(params.Get("endport"))
	if err != nil {
		http.Error(w, "Parsing param error:", http.StatusBadRequest)
		return
	}

	valid := validation(w, hostname, startPort, endPort)
	if !valid {
		return
	}

	ctx := r.Context()
	ports := scanner.ScanRange(ctx, hostname, startPort, endPort)

	response := &ScanResponse{Hostname: hostname, Ports: ports}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON encode error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/scan", handler)

	fmt.Println("Starting web-server on port 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Starting server error:", err)
		return
	}
}
