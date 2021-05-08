package main_test

import (
	"net"
	"testing"
)

func CreateConn(t *testing.T) (net.Conn, net.Conn) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	var client, server net.Conn
	go func() {
		defer ln.Close()
		server, err = ln.Accept()
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
	}()

	client, err = net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	return client, server
}
