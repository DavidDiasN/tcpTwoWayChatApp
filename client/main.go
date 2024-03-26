package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting up client")
	conn, err := net.Dial("tcp", "localhost:5003")
	if err != nil {
		log.Fatal(fmt.Errorf("ran into an error trying to connect to server: %v", err))
	}
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
			fmt.Println("Read called from client")
		}
	}()

	for {

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(message))

		if err != nil {
			log.Fatal(fmt.Errorf("you got an error writing to the server: %v", err))
		}
		fmt.Println("Write called on client")
	}
}
