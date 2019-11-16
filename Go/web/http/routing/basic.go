package main

import (
	"io"
	"log"
	"net/http"
)

const (
	PORT = "8080"
	HOST = "localhost"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello World!")
}

func login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "logged in!")
}

func logout(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "logged out!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
