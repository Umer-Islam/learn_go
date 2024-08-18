package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form error %v", err)
		return
	}
	fmt.Fprintf(w, "print sucessfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "./form" {
		http.Error(w, "404 not found😭 ", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not get ", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello 😀")

}
func main() {

	var fileServer = http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("./form", formHandler)
	http.HandleFunc("./hello", helloHandler)
	fmt.Printf("starting at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
