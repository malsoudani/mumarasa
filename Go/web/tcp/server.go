package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	PORT      = "8080"
	HOST      = "localhost"
	CONN_TYPE = "tcp"
)

// in this example we create a tcp server using the net package that go provides.
// it is better to create a tcp server than an http one when designing highly effecient systems
// lets take it line by line

func main() {
	listener, err := net.Listen(CONN_TYPE, HOST+":"+PORT) // this creates the tcp server and listens on port 8080
	if err != nil {                                       // handle your errors
		log.Fatal("error starting tcp server: ", err)
	}
	defer listener.Close()                           // defer closing the listener for when the program stops other wise you are gonna be hogging that port
	log.Println("listening on " + HOST + ":" + PORT) // log some stuff on the command line
	for {                                            // this is where the magic kinda happens! so this represents our "event" loop in the sense that this is a constant loop that accepts stuff form the listener and logs the response
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println(conn)
	}
}

func handleRequest() {
	msg, err := bufio.NewReader(conn).ReadString("\n")
	if err != nil {
		fmt.Println("Error reading: ", error.Error())
	}
	fmt.Print("message received from the client: ", string(msg))
}
