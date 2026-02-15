package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/qweviluxx/NetViper.git/internal/repository"
	internal "github.com/qweviluxx/NetViper.git/internal/scanner"
)

type ScanRequest struct {
	Hostname  string `json:"hostname"`
	StartPort int    `json:"startport"`
	EndPort   int    `json:"endport"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func validation(h string, s, e int) bool {

	if h == "" || s < 0 || e < 0 || e < s || s > 65535 || e > 65535 {

		return false
	}

	return true
}

func wsHandler(w http.ResponseWriter, r *http.Request, repo repository.Repository, scanner *internal.Scanner) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	var req ScanRequest
	if err := conn.ReadJSON(&req); err != nil {
		return
	}

	if !validation(req.Hostname, req.StartPort, req.EndPort) {
		conn.WriteJSON(map[string]string{"error": "Invalid parameters"})
		return
	}

	out := make(chan int)
	go scanner.ScanRange(r.Context(), req.Hostname, req.StartPort, req.EndPort, out)

	foundPorts := []int{}
	for port := range out {
		foundPorts = append(foundPorts, port)
		conn.WriteJSON(map[string]int{"port": port})
	}

	repo.SaveDB(foundPorts, req.Hostname)
}

func historyHandler(w http.ResponseWriter, r *http.Request, repo repository.Repository) {
	data, err := repo.Receiver()
	if err != nil {
		http.Error(w, "Failed to get history", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(data))
}

func main() {
	scanner := internal.NewScanner("tcp")

	repo, err := repository.New("./scanner.db")
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		return
	}

	http.HandleFunc("/scan", func(w http.ResponseWriter, r *http.Request) { wsHandler(w, r, repo, scanner) })
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) { historyHandler(w, r, repo) })

	fmt.Println("Starting web-server on port 8080...")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Starting server error:", err)
		return
	}

}
