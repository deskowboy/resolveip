package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

func request() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	fn := func(ctx context.Context, network, addr string) (net.Conn, error) {
		fmt.Println("address original =", addr)
		if addr == "www.cskl119.xyz:443" {
			addr = "179.67.163.159:443"
			fmt.Println("address modified =", addr)
		}
		return dialer.DialContext(ctx, network, addr)
	}

	client := req.C() // Use C() to create a client.
	client.SetDial(fn)
	resp, err := client.R(). // Use R() to create a request.
		Get("https://www.cskl119.xyz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func example() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	// or create your own transport, there's an example on godoc.
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		fmt.Println("address original =", addr)
		if addr == "google.com:443" {
			addr = "216.58.198.206:443"
			fmt.Println("address modified =", addr)
		}
		return dialer.DialContext(ctx, network, addr)
	}
	resp, err := http.Get("https://google.com")
	log.Println(resp.Header, err)
}
func main() {
	request()
	//example()
}
