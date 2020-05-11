package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		mess, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Received:", string(mess))

		match, _ := regexp.MatchString("([a-z]+)", mess)
		if !match {
			mess := strings.Trim(mess, "\r\n")
			number, err := strconv.Atoi(mess)
			if err != nil {
				log.Fatal(err)
			}
			conn.Write([]byte(strconv.Itoa(number*2) + "\n"))
		} else {
			conn.Write([]byte(strings.ToUpper(mess) + "\n"))
		}
	}
}
