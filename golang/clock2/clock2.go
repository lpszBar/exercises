package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil { log.Fatal(err) }
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)  // conection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("14:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
