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
	ADMIN_USER     = "admin"     // adding basic auth ... V2
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

// v2 added  BasicAuth function
// lets talk about adding basic auth:
// 1. the first thing to remember from the the first version of this iteration is that http.ListenAndServe expected a handler to be passed in as its second argument which means that it expected a function reference for its arguments
// notice that we are making a handler factory in BasicAuth method which returns a closure "func reference"
// which means that at the time of passing the function down to the ListenAndServe method, nothing about the basic auth mechanism has been invoked, however your hello world handler which evaluates after handling authentication would have been bound to the handler variable but not invoked nor evaluated and same thing goes for the realm variable but thats pretty obvious since its just a string rather than a func.

// 2. BasicAuth is a part of the net/http package, it is invokable only on a type *Request
// if we were to run the following command on the command line while this server is running: curl -i -H 'Authorization:Basic YWRtaW46c2VjcmV0' http://localhost:8080/
// notice that we are passing a basic Auth header with base64 encoded string "admin:secret" and it returned a "Hello World!"
// while it would return "You are unauthorized to access the application." otherwise.
