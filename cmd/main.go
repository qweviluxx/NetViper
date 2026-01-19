package main

import (
	"fmt"
	"time"

	. "github.com/qweviluxx/GopherScanner.git/internal"
)

func main() {
	s := NewScanner("tcp")
	conn, err := s.ScanPort("scanme.nmap.org", 80)
	fmt.Println(conn)
	fmt.Println(err)

	start := time.Now()
	ports := s.ScanRange("scanme.nmap.org", 20, 1000)
	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(ports)
}
