package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"github.com/mhutchinson/insecurity/proxy"
)

var (
	host   = flag.String("host", "example.com:80", "hostname or IP address to proxy to")
	listen = flag.String("listen", ":8080", "interface and port to listen on")
)

func main() {
	flag.Parse()

	p := proxy.NewProxy(*host)

	listener, err := net.Listen("tcp", *listen)
	if err != nil {
		glog.Exitf("Failed to set up listener: %v", err)
	}

	for {
		in, err := listener.Accept()
		if err != nil {
			glog.Errorf("Failed to accept connection: %v", err)
			continue
		}
		glog.Info("Accepted and proxying connection")
		go p.Handle(in)
	}
}
