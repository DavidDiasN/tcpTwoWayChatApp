package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":5003")
	if err != nil {
		log.Fatal(fmt.Errorf("there was an issue making the server: %v", err))
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(fmt.Errorf("there was an error accepting the connection: %v", err))
		}
		fmt.Println("Succesful connection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			buffer := make([]byte, 100)
			n, err := conn.Read(buffer)
			if err != nil {
				continue
			}
			if n == -1 {
				continue
			}
			if string(buffer) == "leave" {
				return
			}
			fmt.Println(string(buffer))
			fmt.Println("Read called from server")
		}
	}()

	for {

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return err
		}

		_, err = conn.Write([]byte(message))

		if err != nil {
			log.Fatal(fmt.Errorf("you got an error writing to the server: %v", err))
		}
		fmt.Println("Write called on server")

	}
}
