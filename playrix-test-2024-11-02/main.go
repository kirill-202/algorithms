package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"gopkg.in/Iwark/spreadsheet.v2"
	"hash/crc32"
	"strings"
	"encoding/json"
	"io"
)


const LOG_FILE_PATH string = "./logs.txt"
const PLAYRIX_SPREAD_SHEET = "1eKxkgpwtTeDSq3R9XHMvVszyyscWp8HJUW9YV1K46es"
const GOOGLE_CRED_PATH string = "./client_secret.json"
const GREEDLY_API_KEY string = "JiuFBPvsVtVFIr"
var GREEDLY_DATABASE_ID string = "v4yj834ypiupv"
var SHEET_NAMES = os.Getenv("SHEET_NAMES")

var eventLogger *log.Logger
var errorLogger *log.Logger
var logMutex sync.Mutex
var InitSheetHash []SheetHash


func LogEvent(eventLogger *log.Logger, message string) {
    logMutex.Lock()
    defer logMutex.Unlock()
    eventLogger.Println(message)
}

type Grid struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Status  string   `json:"status"`
    Columns []Column `json:"columns"`
}

type View struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    GridStatus  string   `json:"gridStatus"`
	GridId  string `json:"gridId"`
    Columns []Column `json:"columns"`
}

type Column struct {
    ID            string       `json:"id"`
    Name          string       `json:"name"`
    Type          string       `json:"type"`
    LanguageCode  string       `json:"languageCode,omitempty"`
    IsSource      bool         `json:"isSource"`
    IsTarget      bool         `json:"isTarget"`
    DependsOn     string       `json:"dependsOn,omitempty"`

}

type Record struct {
    ID    string `json:"id"`
    Cells []Cell `json:"cells"`
}

//Columns are actually rows {column1 None} and so on
type Cell struct {
    //ColumnID string     `json:"columnId"`
    Value    string 	`json:"value"`
}



type RowHash struct {
	RowID int
	RowCheckSum uint32
}

type SheetHash struct {
	SheetID int
	HashedColumns []RowHash
	
}


type GridlyClient struct {
	APIKey     string
	HTTPClient *http.Client
	BaseURL    string
}

func NewGridlyClient(apiKey string) *GridlyClient {
	return &GridlyClient{
		APIKey:     apiKey,
		HTTPClient: http.DefaultClient,
		BaseURL:    "https://api.gridly.com/v1",
	}
}


func (c *GridlyClient) doGETRequest(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "ApiKey "+c.APIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return body, nil
}

// Fetch list of grids
func (c *GridlyClient) GetGreedlyTables(databaseID string) ([]Grid, error) {
	endpoint := fmt.Sprintf("/grids?dbId=%s", databaseID)
	body, err := c.doGETRequest(endpoint)
	if err != nil {
		return nil, err
	}

	var grids []Grid
	if err := json.Unmarshal(body, &grids); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return grids, nil
}

// Fetch views for a specific grid
func (c *GridlyClient) GetView(gridID string) (View, error) {
	endpoint := fmt.Sprintf("/views?gridId=%s", gridID)
	body, err := c.doGETRequest(endpoint)
	if err != nil {
		return View{}, err
	}

	var views []View
	if err := json.Unmarshal(body, &views); err != nil {
		return View{}, fmt.Errorf("error parsing JSON: %v", err)
	}
	if len(views) > 0 {
		return views[0], nil
	}
	return View{}, fmt.Errorf("no views found")
}

// Fetch records for a specific view
func (c *GridlyClient) GetViewRecords(viewID string) ([]Record, error) {
	endpoint := fmt.Sprintf("/views/%s/records", viewID)
	body, err := c.doGETRequest(endpoint)
	if err != nil {
		return nil, err
	}

	var records []Record
	if err := json.Unmarshal(body, &records); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return records, nil
}


func GetSheetList(sheetsString string, separator string) ([]string, error) {
	if sheetsString == "" {
		return nil, fmt.Errorf("no sheets name in enviromant varaibles")
    }
	return strings.Split(sheetsString, separator), nil
}

func ProcessSheet(id int, ss spreadsheet.Spreadsheet, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

    sheet, err := ss.SheetByIndex(uint(id))
    if err != nil {
        logMessage := fmt.Sprintf("can't open sheet %d : %v", id, err)
        LogEvent(errorLogger, logMessage)
        return
    }

    hashedColumns := HashRows(sheet)

    // Lock to prevent concurrent write issues
    mu.Lock()
    InitSheetHash = append(InitSheetHash, SheetHash{SheetID: id, HashedColumns: hashedColumns})
    mu.Unlock()

    logMessage := fmt.Sprintf("all columns of sheet with id %d were hashed", id)
    LogEvent(eventLogger, logMessage)
}

func HashRows(sh *spreadsheet.Sheet) []RowHash {
	/*
	Hash Column values for each Column and return CheckSum with Column ID
	NOTE: it will be better to use rows instewad of columns, but Greedly operates over columns
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
	//fmt.Println(cellValues)
	//cellValues [GameTip_1 John Протираем пыль... Wiping the dust... 50 1 -]
	concatenatedValues := strings.Join(cellValues, "")
	
	table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum([]byte(concatenatedValues), table)
	return checksum
}

func main() {


	fmt.Println("The programm has started...")


	logFile, err := os.OpenFile(LOG_FILE_PATH, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open the log file: %v\n", err)

    }
    defer logFile.Close()


	eventLogger = log.New(logFile, "[EVENT]", log.LstdFlags)
	errorLogger = log.New(logFile, "[ERROR]", log.LstdFlags)

	sheetNames, err := GetSheetList(SHEET_NAMES, ","); if err != nil {
		logText := fmt.Sprintln("Error reading sheet titles from environment", err)
		LogEvent(errorLogger, logText)
	}
	
	LogEvent(eventLogger, fmt.Sprintln("Sheet titles to sync in GoogleSheet and Gridly: ", sheetNames))



	service, err := spreadsheet.NewService(); if err != nil {
		logMessage := fmt.Sprintf("Failed to set up service to work with sheets, check validity of client_secret file: %v", err)
		LogEvent(errorLogger, logMessage)
		os.Exit(1)

		
	}
	
	ssheet, err := service.FetchSpreadsheet(PLAYRIX_SPREAD_SHEET); if err != nil {
		logMessage := fmt.Sprintf("can't open spreedsheet %s : %v", PLAYRIX_SPREAD_SHEET, err)
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


	client := NewGridlyClient(GREEDLY_API_KEY)

	grids, err := client.GetGreedlyTables(GREEDLY_DATABASE_ID)
	if err != nil {
		fmt.Println("Error fetching grids:", err)
	} else {
		fmt.Println("Grids:", grids)
	}
	for _, grid := range grids {
		view, err := client.GetView(grid.ID)
		if err != nil {
			fmt.Println("Error fetching view:", err)
		} else {
			fmt.Println("View:", view)
		}

		records, err := client.GetViewRecords(view.ID)
		if err != nil {
			fmt.Println("Error fetching records:", err)
		} else {
			fmt.Println("Records:", records)
		}
	}




	LogEvent(eventLogger, "All sheets are processed and hashed.")
}