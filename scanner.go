package gopherscanner

import (
	"fmt"
	"net"
	"time"
)

var timeout time.Duration = 1 * time.Second

func ScanPort(protocol, hostname string, port int) (bool, error) {
	address := fmt.Sprintf(hostname, ":%d", port)
	conn, err := net.DialTimeout(protocol, address, timeout)

	if err != nil {
		fmt.Printf("TCP connection error: %w", err)
		return false, err
	}
	conn.Close()
	return true, nil
}
