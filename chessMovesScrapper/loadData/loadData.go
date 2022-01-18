package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create table, if not already there, to store details of moves from data.csv file.
	database, err := sql.Open("sqlite3", "./../moves.db")
	if err != nil {
		log.Fatalf("Error while openeing db %q", err)
		return
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS chessmoves (id VARCHAR(5) PRIMARY KEY, name TEXT, moves TEXT)")
	if err != nil {
		log.Fatalf("Error while creating table %q", err)
		return
	}
	statement.Exec()

	// Start populating the DB
	// While there's still data in data.csv generated using fetchData.go

	// Create a new file to store the data from scrapping
	fileName := "data.csv"
	file, err := os.Open(fileName)

	// If any error occured log it to console.
	if err != nil {
		log.Fatalf("Could not open the file, err: %q\n", err)
		return
	}
	defer file.Close()

	// Create a write to write result to the file.
	reader := csv.NewReader(file)

	statement, _ = database.Prepare("INSERT INTO chessmoves (id, name, moves) VALUES (?, ?, ?)")

	fmt.Println("STARTED INSERTING DATA")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading records from csv file: %q\n", err)
			return
		}

		statement.Exec(record[0], record[1], record[2])
	}

	fmt.Println("DATA INSERTION COMPLETED WITH NO ERRORS")
}
