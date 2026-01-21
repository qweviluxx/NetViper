package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qweviluxx/GopherScanner.git/internal"
)

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

	if hostname == "" {
		http.Error(w, "Parsing param error:", http.StatusBadRequest)
		return
	}

	ports := scanner.ScanRange(hostname, startPort, endPort)
	fmt.Fprintf(w, "opened ports:%v", ports)

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
