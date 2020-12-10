package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/mhutchinson/insecurity/discovery"
)

var (
	host    = flag.String("host", "", "hostname or IP address to scan")
	start   = flag.Int("from", 1, "port number to start scanning from")
	end     = flag.Int("to", 1024, "port number to end scanning at")
	workers = flag.Int("workers", 10, "number of concurrent threads to scan with")
)

func main() {
	flag.Parse()

	open := discovery.Scan(*host, *start, *end, *workers)
	for _, p := range open {
		glog.Infof("%d open", p)
	}
}
