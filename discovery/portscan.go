package discovery

import (
	"fmt"
	"net"
	"sort"

	"github.com/golang/glog"
)

// Scan port scans the range [start, end] on the given host, using the number
// of workers specified. Returs the open ports, sorted in increasing order.
func Scan(host string, start, end, workers int) []int {
	candidates := make(chan int, workers)
	results := make(chan int)
	defer close(candidates)
	defer close(results)

	for i := 0; i < workers; i++ {
		go worker(host, candidates, results)
	}

	go func() {
		for i := start; i <= end; i++ {
			candidates <- i
		}
	}()

	var openPorts []int
	for i := start; i <= end; i++ {
		port := <-results
		if port > 0 {
			openPorts = append(openPorts, port)
		}
	}
	sort.Ints(openPorts)
	return openPorts
}

// worker keeps consuming port numbers from candidates until closed.
// outputs a value to results for every candidate, which will be the
// port number if open, or 0 if not.
func worker(host string, candidates <-chan int, results chan<- int) {
	for p := range candidates {
		address := fmt.Sprintf("%s:%d", host, p)
		glog.V(3).Infof("scanning %s", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}
