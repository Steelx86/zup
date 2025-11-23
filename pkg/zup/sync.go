package zup

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func startTCPServer(certFile string, keyFile string) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Error loading TLS certificate: %v", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	listener, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	fmt.Println("listen on TCP port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close() // closes when finished

	fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from connection: %v", err)
			return
		}

		data := buffer[:n]
		fmt.Printf("Received data from %s: %s\n", conn.RemoteAddr().String(), string(data))

		_, err = conn.Write(data)
		if err != nil {
			log.Printf("Error writing to connection: %v", err)
			return
		}
	}
}
