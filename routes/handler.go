package routes

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is ibcradible")
}

func article_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Articles")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/articles/", article_page)
	http.ListenAndServe(":8080", nil)
}
