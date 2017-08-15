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
	defer li.Close()

	fmt.Println("Listening on Port 8080")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Starting Connection: %s\n", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Printf("Closing Connection: %s\n", conn.RemoteAddr())
}
