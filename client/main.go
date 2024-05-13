package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	cert, err := tls.LoadX509KeyPair("../certs/client.crt", "../certs/client.key")
	if err != nil {
		fmt.Println("Error loading client certificates:", err)
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
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certpool,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "localhost:8888", &config)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server!")

	_, err = conn.Write([]byte("Hello from client!"))
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Println("Received from server:", string(buf[:n]))
}
