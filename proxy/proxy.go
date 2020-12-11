package proxy

import (
	"io"
	"net"

	"github.com/golang/glog"
)

// Proxy is a handler for connections which will transparently be sent
// to a configured address.
type Proxy struct {
	address string
}

// NewProxy creates a Proxy to the given TCP address.
func NewProxy(address string) Proxy {
	return Proxy{
		address: address,
	}
}

// Handle proxies the inbound connection to a new connection that will
// be made to the remote server.
func (p Proxy) Handle(in net.Conn) {
	out, err := net.Dial("tcp", p.address)
	if err != nil {
		glog.Errorf("failed to establish connection to %q", p.address)
		return
	}
	defer out.Close()

	go func() {
		if _, err := io.Copy(out, in); err != nil {
			glog.Errorf("error copying request to remote: %v", err)
		}
	}()

	if _, err := io.Copy(in, out); err != nil {
		glog.Errorf("error copying results to requester: %v", err)
	}
}
