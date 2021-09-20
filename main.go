package main

//49.34
import (
	"fmt"
	"html/template"
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
		<b>Some text</b>
		<ol>
			<li>ferubko</li>
			<li>krylov</li>
			<li>labutin</li>
		</ol>
		<a href="//vk.com/ferubko99/">Создатель сайта</a>
		<a href="#anchor">Перейти к якорю</a>
		<dl>
			<dt>MSU</dt><dd> <li> Chemistry Faculty</li>
				<li>Physical Faculty</li></dd>
			<dt>MPTI</dt><dd> <li> FPMI</li>
				<li> FIPT</li></dd>
		</dl>
		<figure>
 			 <img src="//w-dog.ru/wallpapers/9/3/451292977977142/okean-kamni-plyazh-rassvet-gorizont-solnce.jpg">
  			<figcaption>
    			Институт океанологии
  			</figcaption>
		</figure>
		<h1 id="anchor">Якорь</h1>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, st1)
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
