package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func atm(c net.Conn, name string) {

}

func atmHandler(c net.Conn) {
	fmt.Println("은행을 이용할 사용자 이름을 입력해주세요 : ")
	input := bufio.NewScanner(c)
	for input.Scan() {
		go atm(c, input.Text())
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ATM is running...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go atmHandler(conn)
	}

}
