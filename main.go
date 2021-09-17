package main

import (
	"fmt"
	"net/http"
)

type Station struct {
	name             string
	northernLatitude float64
	longitude        float64
	whyLongit        string
}
type User struct {
	name         string
	email        string
	organisation string
}

func (s *Station) getAllInfo() string {
	return fmt.Sprintf(" station %s is located on %f north latitude ang %f %s longit",
		s.name, s.northernLatitude, s.longitude, s.whyLongit)
}

func (s *Station) setNewMame(newName string) {
	s.name = newName
}
func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{name: "Andrey Ferubko", email: "ferubko1999@gmail.com", organisation: "LaserLab MSU"}
	st1 := Station{name: "5586", northernLatitude: 73.1105, longitude: 61.3195, whyLongit: "east"}
	fmt.Fprintf(w, `<h1>Russia North Ecology Base</h1>
		<b>Some text</b>`)
	fmt.Fprintf(w, "User: "+bob.name+"\n")
	fmt.Fprintf(w, "Station: "+st1.getAllInfo()+"\n")
	st1.setNewMame("Laptev See")
	fmt.Fprintf(w, "Station: "+st1.getAllInfo()+"\n")
}

func article_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Articles set will be here. ")
}

func stations_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The list of stations will be here")
}

func map_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Map will be here")
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/articles/", article_page)
	http.HandleFunc("/stations/", stations_page)
	http.HandleFunc("/map/", map_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
