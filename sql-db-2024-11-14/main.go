package main

import (
	"os"
	"fmt"
	"database/sql"
	"strings"
	"encoding/csv"

	_ "github.com/mattn/go-sqlite3"
)

const DB_NAME string = "main.db"


type Customer struct {
	Name string
	Email string
}


func DBExist(filename string) (bool, error) {
	if !strings.HasSuffix(filename, ".db") {
		return false, fmt.Errorf("file extension doesn't match sqlLite")
	}

	_, err := os.Stat(filename); if err != nil {
		return false, err
	}

	return true, nil
}

func SetDBTables(filename string) error {
	db, err := sql.Open("sqlite3", filename); if err != nil {
		return err
	}
	defer db.Close()

	customers_table_query := `CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT NOT NULL
	);`
	_,  err = db.Exec(customers_table_query); if err != nil {
		return err
	}
	return nil
}


func main() {

	exists, err := DBExist(DB_NAME); if !exists {
		fmt.Println(err)
		fmt.Println("Setting new DB with empty tables..")
		err := SetDBTables(DB_NAME); if err != nil {
			fmt.Println("Can't create tables", err)
			os.Exit(1)
		}
	}

	file, err := os.Open("customers.csv"); if err != nil {
		fmt.Println("Can't read csv file", err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	headers, err := r.Read(); if err != nil {
		fmt.Println("Can't read headers from file", err)
		os.Exit(1)
	}
	fmt.Printf("Headers: %v\n", headers)

	var Customers []Customer
	records, err := r.ReadAll(); if err != nil {
		fmt.Println("Can't read rows from file", err)
		os.Exit(1)
	}
	//fmt.Println(records)
	for _, record := range records {
		Customers = append(Customers, Customer{Name:record[1], Email:record[2]})
	}
	fmt.Println(Customers)


	db, err := sql.Open("sqlite3", DB_NAME); if err != nil {
		fmt.Println("Can't open database", err)
	}
	defer db.Close()


	for _, customer := range Customers {
		query := `INSERT INTO customers (name, email) VALUES (?, ?);`

		_, err := db.Exec(query, customer.Name, customer.Email); if err != nil {
			fmt.Println("Can't write the row for ", customer)
		}
	}

	rows, err := db.Query("SELECT * FROM customers;")
	if err != nil {
		fmt.Println("Can't read from table:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Row error:", err)
	}

}