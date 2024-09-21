package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", handleget)
	http.HandleFunc("/query", handlegetname)
	//http://localhost:8080/query?name=JohnDoe
	//
	port := ":8080"
	fmt.Println("Server is listening on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handlegetname(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	name := params.Get("name")
	fmt.Fprintf(w, "Hi %s ,How are you?", name)
}

//func handleget(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "Server is running ,serve is made by Prashant Pandey !!!")
//}
