package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is ibcradible")
}

func main() {
	fmt.Println("Go рулит")
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)

}
