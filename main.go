package main

import (
	"net/http"
	"net/url"
)



// kelvin to fahrenheit 1.8*(K-273) + 32
type Weatherinfo struct {
	Name : string `json : "name"`
	Temp : string `json: "temp"`
	High : string `json : "temp_max"`
	Low : string `json : "temp_min"`
	Humidity : string `json : "humidity"`
	
}
func Weatherinput(w http.ResponseWriter, r *http.Request) string {
	city := ""

	err := r.ParseForm()
	if err != nil {
		city = err.PostFormValue("name")
	}

	return city
}

// 46074bec0377037004820d9c079cdad9 
func main() {


	key := "46074bec0377037004820d9c079cdad9"
	resp, err := http.PostForm("https://api.openweathermap.org/data/2.5/weather?q={city name}&appid={API key}", url.Values{"city name": {city}, "API key" : {key})
}
