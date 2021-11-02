package main

import (
	"io"
	"log"
	"net"
	"os"
)

func start(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	go start(os.Stdout, conn)
	start(conn, os.Stdin)
}
