package main

//1.10.00
import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"

	_ "github.com/lib/pq"
)

type Station struct {
	Name             string
	NorthernLatitude float64
	Longitude        float64
	WhyLongit        string
	DeepOfSelection  []float64
}
type PostgStations struct {
	Name             string
	NorthernLatitude float64
	Longitude        float64
}
type User struct {
	Name         string
	Email        string
	Organisation string
}

func (s *Station) getAllInfo() string {
	return fmt.Sprintf(" station %s is located on %f north latitude ang %f %s longit",
		s.Name, s.NorthernLatitude, s.Longitude, s.WhyLongit)
}

func (s *Station) setNewMame(newName string) {
	s.Name = newName
}
func home_page(w http.ResponseWriter, r *http.Request) {
	//bob := User{name: "Andrey Ferubko", email: "ferubko1999@gmail.com", organisation: "LaserLab MSU"}
	st1 := Station{Name: "5586", NorthernLatitude: 73.1105, Longitude: 61.3195, WhyLongit: "east",
		DeepOfSelection: []float64{10, 100, 99.1}}
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, st1)
	//fmt.Fprintf(w, "User: "+bob.name+"\n")
	//fmt.Fprintf(w, "Station: "+st1.getAllInfo()+"\n")
	//st1.setNewMame("Laptev See")
	//fmt.Fprintf(w, "Station: "+st1.getAllInfo()+"\n")
}

func expeditions_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/expeditions_page.html")
	tmpl.Execute(w, 122)
}
func transects_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/transects_page.html")
	tmpl.Execute(w, 122)
}
func news_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/news_page.html")
	tmpl.Execute(w, 122)
}
func methods_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/methods_page.html")
	tmpl.Execute(w, 122)
}
func publications_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/publications_page.html")
	tmpl.Execute(w, 122)
}
func map_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/mapfirst.html")
	tmpl.Execute(w, 122)
}
func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/map/", map_page)
	http.HandleFunc("/expeditions/", expeditions_page)
	http.HandleFunc("/transects/", transects_page)
	http.HandleFunc("/news/", news_page)
	http.HandleFunc("/methods/", methods_page)
	http.HandleFunc("/publications/", publications_page)
	http.ListenAndServe(":8080", nil)
}
func firstpostgr() error {
	connStr := "user=postgres dbname=firstf password=postgres host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	//insert, err := db.Query(`INSERT INTO "Stations" VALUES ('5556',18.9, 90.99)`)
	//(Name, NorthernLatitude, Longitude)
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()
	res, err := db.Query(`SELECT * FROM "public"."Stations"`)
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var object PostgStations
		err := res.Scan(&object.Name, &object.NorthernLatitude, &object.Longitude)
		if err != nil {
			panic(err)
		}
		fmt.Printf(fmt.Sprintf("Station %s on %f on %f", object.Name, object.NorthernLatitude, object.Longitude))
	}
	defer db.Close()
	fmt.Println("bj")
	return nil
}
func excelread(out io.Writer, namefile string, pageex string) error {
	name := "./" + namefile
	xlsx, err := excelize.OpenFile(name)
	if err != nil {
		panic(err)
	}
	rows := xlsx.GetRows(pageex)
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Fprintf(out, "\t", colCell)
		}
		fmt.Fprintln(out)
	}

	return nil
}

func main() {
	//handleRequest()
	excelread(os.Stdout, "Stations.xlsx", "Sheet1")
	err := firstpostgr()
	if err != nil {
		panic(err)
	}
}
