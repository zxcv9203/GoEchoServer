package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleFT(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "42\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	port := "8000"
	listen, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleFT(conn)
	}
}
