package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/brodiep21/weatherapp/weather"
)

var newtemp = template.Must(template.ParseFiles("website/appface.html"))

func Handler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, r.URL.Query()); err != nil {
			http.Error(w, fmt.Sprintf("error handling created template %s", err), http.StatusInternalServerError)
		}

	})
}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	s, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	params := s.Query()
	searchQuery := params.Get("city")

	fmt.Println("Search Query is: ", searchQuery)
}

// func HTMLresponse(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatalf("could not gather form info %s", err)
// 	}

// 	City = r.PostFormValue("city")
// 	// fmt.Printf(w, City)
// }

func main() {
	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", Handler(newtemp)); err != nil {
		log.Fatalf("Could not run server %s", err)
	}

	weatherinputstring := "https://api.openweathermap.org/data/2.5/weather?q=" + +"&appid=46074bec0377037004820d9c079cdad9&units=imperial"
	// key := "46074bec0377037004820d9c079cdad9"
	weather.Weatherinput(weatherinputstring)

}

//
