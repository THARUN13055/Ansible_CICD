package main

import (
	"fmt"
	"log"
	"net/http"
)

func geturl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is the website")
}

func main() {
	http.HandleFunc("/hello", geturl)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
