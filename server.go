package main

import (
	"log"
	"net"
	"time"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":3333")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	log.Printf("New connection from %s", conn.RemoteAddr())

	readBuf := make([]byte, 4096, 4096)

	for {
		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
		readLen, err := conn.Read(readBuf)
		if err != nil {
			log.Printf("From: %s Closed %s", conn.RemoteAddr(), err)
			return
		}
		log.Printf("From: %s Recv: %#v", conn.RemoteAddr(), readBuf[:readLen])
	}
}
