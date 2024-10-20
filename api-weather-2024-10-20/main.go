package main

import (
	"io"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"net/url"
)

const APIKEY string = "bdddf8e187bbf40c1695580de99c24d6"
//https://api.openweathermap.org/data/2.5/weather?q={city name}&appid={API key}
const ENDPOINT string = "https://api.openweathermap.org/data/2.5/weather"


type Weather struct {
	Description string `json:"description"`
	Temperature float64 `json:"temp"`
}

type ApiResponse struct {
	Main    struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}


func main() {

	logger := log.Default()
	logger.SetPrefix("[Weather] ")
	logger.Println("The program has started...")

	var city string


	fmt.Print("Provide a city to check current weather: ")
	_, err := fmt.Scan(&city); if err != nil {
		logger.Fatalln("Can't scan the input:", err)
	}

	u, err  := url.Parse(ENDPOINT); if err !=nil {
		logger.Fatalln("Can't parse url string:", err)
	}


	params := url.Values{}
	params.Set("q", city)
	params.Set("appid", APIKEY)
	u.RawQuery = params.Encode()


	resp, err := http.Get(u.String()); if err != nil {
		logger.Fatalln("Error requesting weather", err, resp.StatusCode)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Fatalln("Request failed with status code:", resp.StatusCode)
	}


	var ar ApiResponse
	bodyBites, err := io.ReadAll(resp.Body); if err != nil {
		logger.Fatalln("Can't read body", err)
	}
	err = json.Unmarshal(bodyBites, &ar); if err != nil {
		logger.Fatalln("Can't unmarshal bites", err)
	}

	weatherData := Weather{Description: ar.Weather[0].Description, Temperature: ar.Main.Temperature}

	fmt.Printf("City -- %s\nWeather -- %s\nTemperature -- %0.2f\nCurrentTime -- %s\n", 
				city, 
				weatherData.Description, 
				weatherData.Temperature, 
				time.Now().Format("2006-01-02T15:04:05"),
			)


}