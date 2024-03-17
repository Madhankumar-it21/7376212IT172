package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/madhan", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Madhankumar")
	})
http.HandleFunc("/rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212IT172")
	})
	fmt.Printf("Server running (port=8081), route: http://localhost:8081/helloworld\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}



