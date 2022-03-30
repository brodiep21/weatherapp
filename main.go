package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Main struct {
	Temp     float64 `json:"temp"`
	High     float64 `json:"temp_max"`
	Low      float64 `json:"temp_min"`
	Humidity int     `json:"humidity"`
}

// kelvin to fahrenheit 1.8*(K-273) + 32
type Weatherinfo struct {
	Main Main
}

func Handler(t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := t.Execute(w, r.URL.Query()); err != nil {
			http.Error(w, fmt.Sprintf("error handling created template %s", err), http.StatusInternalServerError)
		}

	})
}

func Weatherinput(url string) {

	var s Weatherinfo

	client := &http.Client{Timeout: 3 * time.Second}

	req, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &s)
	// temp := math.Round((s.Main.Temp-273)*1.8 + 32)
	fmt.Println(s)
}

func main() {
	var s string
	fmt.Scanf("%c", &s)

	weatherinputstring := "https://api.openweathermap.org/data/2.5/weather?q=" + s + "&appid=46074bec0377037004820d9c079cdad9&units=imperial"
	// key := "46074bec0377037004820d9c079cdad9"
	Weatherinput(weatherinputstring)

	newtemp := template.Must(template.New("site.html").ParseGlob("website/*.html"))
}

//
