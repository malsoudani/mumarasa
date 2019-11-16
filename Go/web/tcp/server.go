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
	log.Println("Listening on " + HOST + ":" + PORT) // log some stuff on the command line
	for {                                            // this is where the magic kinda happens! so this represents our "event" loop in the sense that this is a constant loop that accepts stuff form the listener and logs the response
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go handleRequest(conn) // V2 now we handle the connection within a custom handler and invoking it as a goroutine
	}
}

func handleRequest(conn net.Conn) { // V2 is to handle the request that was sniffed by listener.Accept(), we are writing a custom handler for that purpose called handleRequest(), handleRequest reads the connection into the buffer "bufio" until the first occurence of "\n", handles errors, and if all goes well it prints it to the standard output
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Print("message received from the client: ", string(msg))
	conn.Write([]byte(msg)) // V3: writing back to the connection stream, now when running your tcp server you can write back the client this way
	conn.Close()
}
