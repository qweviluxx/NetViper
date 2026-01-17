package main

import (
	"fmt"

	"github.com/qweviluxx/GopherScanner.git/internal"
)

func main() {
	conn, err := internal.ScanPort("tcp", "scanme.nmap.org", 80)
	fmt.Println(conn)
	fmt.Println(err)
}
