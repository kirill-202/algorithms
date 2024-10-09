package main

import (
	"fmt"
	"regexp"
	"strconv"
	"os"
)


type Weekday int 

type Date struct {
	Year int
	Month int
	Day int
	Weekday Weekday
}

func FromString(dateString string) (*Date, error) {
	datePattern := `^(\d{4})-(\d{2})-(\d{2})$`
	re := regexp.MustCompile(datePattern)

	matches := re.FindStringSubmatch(dateString)
    if len(matches) != 4 {
        return nil, fmt.Errorf("invalid date format: %s, use the layout yyyy-mm-dd",dateString)
    }
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])

	date := &Date{Year: year, Month: month, Day: day}
    date.calculateWeekday() 

	return date, nil
}

func (d *Date) calculateWeekday() {
    month:= d.Month
   year:= d.Year

    // Zeller's Congruence requires January and February to be counted as months 13 and 14 of the previous year
    if month< 3 {
        month+= 12
        year--
    }

    centYear :=year% 100        
    century :=year/ 100         

    // Zeller's formula
    f := d.Day + (13*(month+ 1))/5 + centYear + (centYear / 4) + (century / 4) - (2* century)

    // Calculate the weekday
    d.Weekday = Weekday((f % 7 + 7) % 7) // To ensure non-negative value
	d.Weekday = (d.Weekday + 6) % 7
}


const (
    Sunday Weekday = iota // Sunday == 0
    Monday                // Monday == 1
    Tuesday               // Tuesday == 2
    Wednesday             // Wednesday == 3
    Thursday              // Thursday == 4
    Friday                // Friday == 5
    Saturday              // Saturday == 6
)


func (d Weekday) String() string {
    dayMap := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    return dayMap[d]
}

var UnixStartDate = Date{
	Year : 1970,
	Month : 1,
	Day : 1,
	Weekday: 4,
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide the date in the format yyyy-mm-dd")
		os.Exit(1)
	}

	myDate, err := FromString(os.Args[1]); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Weekday", myDate.Weekday)
}



