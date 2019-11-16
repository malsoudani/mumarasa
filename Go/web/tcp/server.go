package main

import (
	"log"
	"net"
)

const (
	PORT      = "8080"
	HOST      = "localhost"
	CONN_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONN_TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal("error starting tcp server: ", err)
	}
	defer listener.Close()
	log.Println("listening on " + HOST + ":" + PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println(conn)
	}
}
