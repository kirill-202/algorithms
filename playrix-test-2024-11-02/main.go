package main

import (
	"fmt"
	"log"
	//"net/http"
	"os"
	"sync"
	"gopkg.in/Iwark/spreadsheet.v2"
	"hash/crc32"
	"strings"
)


const LogFilePath string = "./logs.txt"
const PlayRixSpreaddsheet = "1eKxkgpwtTeDSq3R9XHMvVszyyscWp8HJUW9YV1K46es"
const StaticSheetID uint = 0
const GameSheetID uint = 128432535
const GoogleCredsPath string = "./client_secret.json"

var eventLogger *log.Logger
var errorLogger *log.Logger
var logMutex sync.Mutex
var lastSheetHash []SheetHash

func LogEvent(eventLogger *log.Logger, message string) {
    logMutex.Lock()
    defer logMutex.Unlock()
    eventLogger.Println(message)
}

type RowHash struct {
	RowID int
	RowCheckSum uint32
}

type SheetHash struct {
	SheetID int
	HashedRows []RowHash
}

func ProcessSheet(id int, ss spreadsheet.Spreadsheet, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

    sheet, err := ss.SheetByIndex(uint(id))
    if err != nil {
        logMessage := fmt.Sprintf("can't open sheet %d : %v", id, err)
        LogEvent(errorLogger, logMessage)
        return
    }

    hashedRows := HashRows(sheet)

    // Lock to prevent concurrent write issues
    mu.Lock()
    lastSheetHash = append(lastSheetHash, SheetHash{SheetID: id, HashedRows: hashedRows})
    mu.Unlock()

    logMessage := fmt.Sprintf("all rows of sheet with id %d were hashed", id)
    LogEvent(eventLogger, logMessage)
}

func HashRows(sh *spreadsheet.Sheet) []RowHash {
	/*
	Hash Row values for each Row and return CheckSum with Row ID
	*/
	rowHahes := make([]RowHash, 0, 10)

	for index, row := range sh.Rows {
		rowCheckSum := hashRowCells(row)
		rowHahes = append(rowHahes, RowHash{RowID: index, RowCheckSum: rowCheckSum})
	}

	return rowHahes
}

func hashRowCells(rowCells []spreadsheet.Cell) uint32 {
	/* 
	Create a CRC32 checksum table
	Calculate the checksum for the entire concatenated string
	*/
	cellValues := make([]string, 0, 10)
	for _, cell := range rowCells {
		cellValues = append(cellValues, cell.Value)
	}
	concatenatedValues := strings.Join(cellValues, "")
	
	table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum([]byte(concatenatedValues), table)
	return checksum
}

func main() {

	fmt.Println("The programm has started...")

	logFile, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open the log file: %v\n", err)

    }
    defer logFile.Close()


	eventLogger = log.New(logFile, "[EVENT]", log.LstdFlags)
	errorLogger = log.New(logFile, "[ERROR]", log.LstdFlags)

	service, err := spreadsheet.NewService(); if err != nil {
		logMessage := fmt.Sprintf("Failed to set up service to work with sheets, check validity of client_secret file: %v", err)
		LogEvent(errorLogger, logMessage)
		os.Exit(1)

		
	}
	
	ssheet, err := service.FetchSpreadsheet(PlayRixSpreaddsheet); if err != nil {
		logMessage := fmt.Sprintf("can't open spreedsheet %s : %v", PlayRixSpreaddsheet, err)
		LogEvent(errorLogger, logMessage)
		os.Exit(1)
	}

	var wg sync.WaitGroup
    var mu sync.Mutex
	workerNmubers := len(ssheet.Sheets)

	for i:=0; i<workerNmubers; i++ {
		wg.Add(1)
		go ProcessSheet(i, ssheet, &wg, &mu)
	}

	wg.Wait()

	LogEvent(eventLogger, "All sheets are processed and hashed.")
}