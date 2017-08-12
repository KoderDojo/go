package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Listening on port 8080...")

	for {
		fmt.Println("Waiting for connection...")

		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("Connection accepted from %s.\n", conn.RemoteAddr())

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("I heard you say: " + scanner.Text())
	}

	fmt.Printf("Connection terminated by %s.\n", conn.RemoteAddr())
}
