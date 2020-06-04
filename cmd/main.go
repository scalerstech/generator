package main

import (
	"net"
	"fmt"
	"os"
	"math/rand"
	"encoding/base64"
	"time"
)

func main() {
	addr := "0.0.0.0:6502"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen on %v: %v\n", addr, err)
		os.Exit(-1)
	}

	defer l.Close()

	fmt.Printf("listening on %v\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("failed to accept connection: %v\n", err)
		} else {
			fmt.Printf("handling connection: %v\n", conn)
			go handleClient(conn)
		}
	}
}

func handleClient(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		rand.Read(data)
		b64 := base64.StdEncoding.EncodeToString(data)
		conn.Write([]byte(b64))
		time.Sleep(100 * time.Millisecond)
	}
}
