package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../certs/server.crt", "../certs/server.key")
	if err != nil {
		fmt.Println("Error loading server certificates:", err)
		return
	}

	// Load CA certificate
	caCert, err := os.ReadFile("../certs/ca.crt")
	if err != nil {
		fmt.Println("Error loading CA certificate:", err)
		return
	}
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(caCert)

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certpool,
	}

	listener, err := tls.Listen("tcp", "localhost:8888", &config)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error closing listener:", err)
		}
	}(listener)

	fmt.Println("Server listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				fmt.Println("Error closing connection:", err)
			}
		}(conn)

		fmt.Println("Client connected:", conn.RemoteAddr())

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}
	fmt.Println("Received from client:", string(buf[:n]))

	_, err = conn.Write([]byte("Hello from server!"))
	if err != nil {
		fmt.Println("Error writing to client:", err)
		return
	}
}
