package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"

	"gopkg.in/Iwark/spreadsheet.v2"
)


const LogFilePath string = "./logs.txt"
const PlayRixSpreaddsheet = "1eKxkgpwtTeDSq3R9XHMvVszyyscWp8HJUW9YV1K46es"
const StaticSheetID uint = 0
const GameSheetID uint = 128432535
const GoogleCredsPath string = "./client_secret.json"


func main() {

	fmt.Println("The programm has started...")

	logFile, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open file: %v\n", err)
        return
    }
    defer logFile.Close()

	//eventLogger := log.New(logFile, "[EVENT]", log.LstdFlags)

	service, err := spreadsheet.NewService(); if err != nil {
		log.Fatalf("Failed to set up service to work with sheets, check validity of client_secret file: %v\n", err)
	}

	spreedsheet, err := service.FetchSpreadsheet(PlayRixSpreaddsheet); if err != nil {
		log.Fatalf("Can't open spreedsheet %v\n", err)
	}

	gameSheet, err :=  spreedsheet.SheetByID(GameSheetID); if err != nil {
		log.Fatalf("Failed to fetch Game Sheet %v\n", err)
	}

	fmt.Println("Sheet info via rows", gameSheet.Rows[0][1].Value)
	
	staticSheet, err := spreedsheet.SheetByID(StaticSheetID); if err != nil {
		log.Fatalf("Failed to fectch Static Sheet %v\n", err)
	}
	
	fmt.Println("Sheet info via columns", staticSheet.Columns[0][1].Value)



	
}