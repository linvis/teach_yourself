package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go boardcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func boardcast() {
	clients := make(map[client]bool)
	select {
	case cli := <-entering:
		fmt.Printf("1111111")
		clients[cli] = true
	case cli := <-leaving:
		fmt.Printf("2221111111")
		delete(clients, cli)
		close(cli)
	case msg := <-messages:
		for cli := range clients {
			cli <- msg
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "you are " + who
	messages <- who + "has arrived"
	fmt.Print("fsfs")
	entering <- ch
	fmt.Print("222fsfs")

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " " + input.Text()
	}

	leaving <- ch
	messages <- who + "has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
