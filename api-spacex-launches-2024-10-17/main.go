package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"encoding/json"
	"os"
	"strings"
)


type LaunchAPIResponse struct {
	RocketID string `json:"rocket"`
	RawDate string `json:"date_utc"`
	LaunchName string `json:"name"`
	LaunchpadID string `json:"launchpad"`
}

type Rocket struct {
	ID string 
	Name string `json:"name"`
}

type Launch struct {
	Date *time.Time
	Name string
	Location *Location
	Rocket *Rocket
}

type Location struct {
	Name string `json:"locality"`
}


func GetLocation(launchpadID string) (*Location, error) {

	resp, err := http.Get("https://api.spacexdata.com/v4/launchpads/" + launchpadID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Unmarshal into a temporary struct that contains the "location" field
	var location Location

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

func GetRocket(rocketID string) (*Rocket, error) {
	resp, err := http.Get("https://api.spacexdata.com/v4/rockets/" + rocketID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body); if err != nil {
		return nil, err
	}

	rocket := Rocket{ID:rocketID}
	err = json.Unmarshal(body, &rocket); if err != nil {
		return nil, err
	}

	return &rocket, nil
	
}



func NewLaunch(rawResponse LaunchAPIResponse) Launch {

	date, err := ParseDate(rawResponse.RawDate); if err != nil {
		date = nil
	}
	rocket, err := GetRocket(rawResponse.RocketID); if err != nil {
		rocket = nil
	}
	location, err := GetLocation(rawResponse.LaunchpadID); if err != nil {
		location = nil
	}

	return Launch{
		Date: date,
		Name: rawResponse.LaunchName,
		Location: location,
		Rocket: rocket,
	}
	
}


func ParseDate(date string) (*time.Time, error) {
	layout := "2006-01-02T15:04:05.000Z"

	parsedDate, err := time.Parse(layout, date); if err != nil {
		return nil, err
	}
	return &parsedDate, nil
}

func PaginateLaunches(launches []Launch, pageSize int, past bool) {
    var startText string
    if past {
        startText = "Previous"
    } else {
        startText = "Upcoming"
    }

    totalLaunches := len(launches)

	// Calculate the number of pages
    totalPages := (totalLaunches + pageSize - 1) / pageSize  

    for currentPage := 0; currentPage < totalPages; currentPage++ {
        start := currentPage * pageSize
        end := start + pageSize
        if end > totalLaunches {
            end = totalLaunches
        }
        fmt.Printf("\n%s SpaceX Launches (Page %d of %d):\n\n", startText, currentPage+1, totalPages)
        fmt.Printf("%-35s | %-20s | %-15s | %-20s\n", "Name", "Date", "Rocket", "Location")
        fmt.Println(strings.Repeat("-", 90))

        for i, l := range launches[start:end] {
            dateStr := "N/A"
            if l.Date != nil {
                dateStr = l.Date.Format("January 2, 2006")
            }
            rocketName := "N/A"
            if l.Rocket != nil {
                rocketName = l.Rocket.Name
            }
            locationName := "N/A"
            if l.Location != nil {
                locationName = l.Location.Name
            }

            fmt.Printf("%-5d %-30s | %-20s | %-15s | %-20s\n",
                i+1+start, l.Name, dateStr, rocketName, locationName)
        }

        if currentPage < totalPages-1 {
            var input string
            fmt.Print("\nDo you want to see the next page? (y/n): ")
            fmt.Scan(&input)
            if strings.ToLower(input) != "y" {
                break
            }
        }
    }
}

func main() {

	var flag string
	var endpoint string
	var past bool

	fmt.Print("Program has started... Provide the flag: ")

	_, err := fmt.Scan(&flag); if err != nil {
		fmt.Println("Issue with scanning input", err)
		os.Exit(1)
	}

	switch flag {
		case "--past": 
			fmt.Println("Execute past launch logic")
			endpoint = "past"
			past = true

		case "--upcoming":
			fmt.Println("Execute upcoming launch logic")
			endpoint = "upcoming"
			past = false

		default:
			fmt.Println("Incorrect input. Usage: Provide either flag <--upcoming> or <--past>")
			os.Exit(1)
	}

	resp, err := http.Get("https://api.spacexdata.com/v4/launches/" + endpoint)
	if err != nil {
		fmt.Printf("Can't get %s event data: %v\n", endpoint, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}
	

	var responseData []LaunchAPIResponse
	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
		os.Exit(1)
	}

	fmt.Println("Processing response to launch data...")
	launches := make([]Launch, 0, len(responseData))
	for _, v := range responseData {
		launch := NewLaunch(v)
		launches = append(launches, launch)
	}


	PaginateLaunches(launches, 5, past)

}