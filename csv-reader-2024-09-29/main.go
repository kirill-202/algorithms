package main


import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)
const (
    indexName      = 0
    indexAge       = 1
    indexOccupation = 2
)

/*
Use the encoding/csv package to read the contents of a CSV file.
The CSV will contain a list of people with their names, ages, and occupations. Example format:
Copy code
Name, Age, Occupation
John Doe, 29, Engineer
Jane Smith, 34, Doctor
Your program should:
Calculate the average age of all people in the CSV file.
Print the number of people per occupation.
Handle potential errors (e.g., file not found, incorrect format) gracefully, with user-friendly error messages.

*/


func main() {

	peoplePerOccupation := make(map[string]int)
	var totalAge int
	var recordsLength int


	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <csv-file-path>")
		return
	}

	csvFile, err := os.Open(os.Args[1]); if err != nil {
		fmt.Println("error reading file:", os.Args[1])
		return
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)


    if _, err = csvReader.Read(); err != nil {
        fmt.Println("error reading header:", err)
        return
    }

    for {
        record, err := csvReader.Read()
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Printf("Error reading record: %v\n", err)
            return
        }


        if len(record) <= indexOccupation {
            fmt.Println("Invalid record format, skipping...")
            continue
        }


        occupation := strings.TrimSpace(record[indexOccupation])
        peoplePerOccupation[occupation]++

        age, err := strconv.Atoi(strings.TrimSpace(record[indexAge]))
        if err != nil {
            fmt.Printf("Error converting age for record '%v': %v\n", record, err)
            return
        }

        totalAge += age
        recordsLength++
    }

	fmt.Printf("the average is %0.2f \npeople per profession %v\n", float64(totalAge)/float64(recordsLength), peoplePerOccupation)



}