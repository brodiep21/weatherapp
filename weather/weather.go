package weather

import (
	"encoding/json"
	"fmt"
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

// kelvin to fahrenheit 1.8*(K-273) + 32
type Weatherinfo struct {
	Main Main
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