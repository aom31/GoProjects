package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(resp http.ResponseWriter, requ *http.Request) {
	if err := requ.ParseForm(); err != nil {
		fmt.Fprintf(resp, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(resp, "POST request success \n")
	name := requ.FormValue("name")
	email := requ.FormValue("email")
	fmt.Fprintf(resp, "Name = %s \n", name)
	fmt.Fprintf(resp, "Address = %s \n", email)

}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found ", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "method not supported ", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "hello!")
}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
