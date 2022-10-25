package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	connectionName := conn.RemoteAddr().String()

	fmt.Println("New connection", connectionName)
	conn.Write([]byte("Hello" + connectionName))

	defer conn.Close()

	scan := bufio.NewScanner(conn)

	for scan.Scan() {
		text := scan.Text()
		if text == "Exit" {
			conn.Write([]byte("You wanted go out? Okay"))
			fmt.Println("Disconnected" + connectionName)
			break
		} else {
			fmt.Println(connectionName + "enter" + text)
			conn.Write([]byte("You enter" + scan.Text()))

		}
	}
}

func main() {
	listener, _ := net.Listen("tcp", "8080")

	for {
		conn, _ := listener.Accept()

		go handleConnection(conn)
	}

}
