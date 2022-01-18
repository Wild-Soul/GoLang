package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	// Create a new file to store the data from scrapping
	fileName := "./../loadData/data.csv"
	file, err := os.Create(fileName)

	// If any error occured log it to console.
	if err != nil {
		log.Fatalf("Could not create the file, err: %q\n", err)
		return
	}
	defer file.Close()

	// Create a write to write result to the file.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("www.chessgames.com"),
	)

	// Extract information from table element.
	c.OnHTML("table", func(e *colly.HTMLElement) {

		// For each row in table take out information related to ECO code and move name.
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {

			// fmt.Println(el.ChildText("td:nth-child(1)"))
			moveCode := strings.Trim(el.ChildText("td:nth-child(1)"), " ")
			// ecoCode, moveName := details[0], details[1]

			details := el.ChildText("td:nth-child(2)")
			nameAndMoves := strings.Split(details, "\n")

			name, moves := nameAndMoves[0], nameAndMoves[1]

			// Write the details of each move to the .csv file
			writer.Write([]string{moveCode, name, moves})

			// fmt.Println("Details: ", moveCode)
			// fmt.Println("Moves: ", nameAndMoves[0])
		})
	})

	c.Visit("https://www.chessgames.com/chessecohelp.html")

}
