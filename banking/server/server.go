package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func deposit() {

}

func Withdrawal() {

}

func atm(c net.Conn, money string) {

}

func atmHandler(c net.Conn, name string) {
	money := 0
	input := bufio.NewScanner(c)
	for input.Scan() && input.Text() != "exit" {
		fmt.Fprintln(c, name, "님의 현재 잔액 : ", money)
		go atm(c, input.Text())

	}
}

func userHandler(c net.Conn) {
	fmt.Fprint(c, "은행을 이용할 사용자 이름을 입력해주세요 : ")
	input := bufio.NewScanner(c)
	for input.Scan() && input.Text() != "exit" {
		atmHandler(c, input.Text())
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
		go userHandler(conn)
	}

}
