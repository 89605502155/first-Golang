package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Russia North Ecology Map")
}

func article_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Articles set will be here. ")
}

func map_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Map will be here")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/articles/", article_page)
	http.HandleFunc("/map/", map_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
