package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, _ := net.Dial("tcp", ":8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text: ")
		text, _ := reader.ReadString('\n')
		if text == "exit\r\n" {
			os.Exit(0)
		}
		fmt.Fprintf(conn, text+"\n")
		mess, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + mess)
	}
}
