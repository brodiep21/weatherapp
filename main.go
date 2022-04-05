package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

type Main struct {
	Temp     float64 `json:"temp"`
	High     float64 `json:"temp_max"`
	Low      float64 `json:"temp_min"`
	Humidity int     `json:"humidity"`
}

type Weatherinfo struct {
	Main Main
}

var s Weatherinfo

//create an HTML Template
var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("website/*.html"))
}

//string input set for the city
var City string

//receives form method from HTML and parses it into the weather API. Then parses that into a new form template and redirects
func HTMLresponse(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	City = r.FormValue("city")

	weatherinputstring := "https://api.openweathermap.org/data/2.5/weather?q=" + City + "&appid=46074bec0377037004820d9c079cdad9&units=imperial"

	client := &http.Client{Timeout: 3 * time.Second}

	req, err := client.Get(weatherinputstring)
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &s)

	templ.ExecuteTemplate(w, "weather.html", s)
}

//base HTML page for search
func homePage(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "homepage.html", nil)
}

func main() {

	fmt.Println("Starting server at port 8080")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/weather", HTMLresponse)
	http.ListenAndServe(":8080", nil)

}

// key := "46074bec0377037004820d9c079cdad9"
