package main

import (
	"github.com/gorilla/handlers"
	"io"
	"log"
	"net/http"
)

const (
	HOST = "localhost"
	PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(HOST+":"+PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}

// by sending gzipped responses we save network bandwidth and download time eventually rendering the page faster.
// GZIP compression happens as follows:
// the browser starts by sending a request header that tells the server that the browser can accept compressed content (.gzip and .deflate) specifically and if the server can respond with it then it will by sending back the compressed form along with "Content-Encoding : gzip". The point is that the browser only makes a request for a gzipped content and doesn't demand it per se.

// lets talk a little about the http.NewServeMux() method ... well it is pretty self explainatory .. it returns a new ServeMux struct for us. But the more interesting question is what is a ServeMux? well a ServeMux is a `struct` wrapping type that basically represents the router portion of the net/http package. It maps a handler to a distict pattern of a route.
