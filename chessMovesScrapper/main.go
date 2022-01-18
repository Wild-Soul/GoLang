package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Global variable to store connectin to DB
var database *sql.DB

func getAllMovesHelper(w http.ResponseWriter, r *http.Request) {
	// Prepare query to get the code details.
	moveDetails, err := database.Query("SELECT * FROM chessmoves")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	var id, name, moves string
	for moveDetails.Next() {
		moveDetails.Scan(&id, &name, &moves)
		fmt.Fprintf(w, "Code: %v \t Name: %v \t Moves: %v\n", id, name, moves)
		// fmt.Println(id + ": " + name + " " + moves)
	}
}

func getMoveHelper(w http.ResponseWriter, r *http.Request) {
	// Get the params from route
	vars := mux.Vars(r)
	code := vars["code"]
	// fmt.Println(reflect.TypeOf(vars))

	// Prepare query to get the code details.
	moveDetails, err := database.Query("SELECT * FROM chessmoves WHERE id=(?)", code)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var id, name, moves string
	for moveDetails.Next() {
		moveDetails.Scan(&id, &name, &moves)
		fmt.Fprintf(w, "Code: %v \t Name: %v \t Moves: %v\n", id, name, moves)
	}
}

func main() {

	var err error

	// Open database
	database, err = sql.Open("sqlite3", "./moves.db")
	if err != nil {
		log.Fatalf("Error while openeing db %q", err)
		return
	}

	// Create table, if not already there, to store details of moves from data.csv file.
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS chessmoves (id VARCHAR(5) PRIMARY KEY, name TEXT, moves TEXT)")
	if err != nil {
		log.Fatalf("Error while creating table %q", err)
		return
	}
	statement.Exec()

	// Create http server and start serving data.
	r := mux.NewRouter()
	r.HandleFunc("/", getAllMovesHelper)
	r.HandleFunc("/{code}", getMoveHelper)

	http.ListenAndServe(":8090", r)
}
