package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	CONN_HOST      = "localhost" // host
	CONN_PORT      = "8080"      // connection port
	ADMIN_USER     = "admin"     // adding basic auth
	ADMIN_PASSWORD = "secret"
)

func helloWorld(w http.ResponseWriter, r *http.Request) { // a handler generally takes in a request and a resource to write a response with
	fmt.Fprintf(w, "Hello World!") // the response writer works on the HTTP response stream
}

func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are uauthorized to access the application. \n"))
			return
		}
		handler(w, r)
	}

}

func main() {
	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and password")) // register helloWorld as the handler for the "/" path
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)                               // serve requests that handle each http connection on its own goroutine

	if err != nil { // the best way to handle your errors
		log.Fatal("error starting http server : ", err)
		return // if there is an error just exit
	}
}

// this demonstration should print "Hello World!" on port 8080
