package internal

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

type Scanner struct {
	Protocol    string
	timeout     time.Duration
	workerCount int
}

func NewScanner(protocol string) *Scanner {
	return &Scanner{
		Protocol:    protocol,
		timeout:     500 * time.Millisecond,
		workerCount: 100,
	}
}

func (s *Scanner) ScanPort(ctx context.Context, hostname string, port int) (bool, error) {
	address := hostname + fmt.Sprintf(":%d", port)
	d := net.Dialer{}
	childCtx, cancel := context.WithTimeout(ctx, s.timeout)

	conn, err := d.DialContext(childCtx, s.Protocol, address)
	cancel()

	if err != nil {
		return false, err
	}
	conn.Close()
	return true, nil
}

func (s *Scanner) ScanRange(ctx context.Context, hostname string, startPort, endPort int) []int {

	size := endPort - startPort + 1
	openedPorts := []int{}

	ports := make(chan int, size)
	result := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < s.workerCount; i++ {
		wg.Add(1)
		go s.worker(ctx, &wg, hostname, ports, result)
	}

	go func() {
		for j := startPort; j <= endPort; j++ {
			ports <- j
		}
		close(ports)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for p := range result {
		res := p
		if res != 0 {
			openedPorts = append(openedPorts, res)
		}
	}

	return openedPorts
}

func (s *Scanner) worker(ctx context.Context, wg *sync.WaitGroup, hostname string, ports chan int, result chan int) {
	defer wg.Done()
	for i := range ports {
		select {
		case <-ctx.Done():
			return
		default:
		}
		ok, _ := s.ScanPort(ctx, hostname, i)
		if ok {
			result <- i
		} else {
			result <- 0
		}
	}
}
