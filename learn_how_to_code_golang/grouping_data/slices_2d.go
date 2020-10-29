package main

import "fmt"

func main() {

	james := []string{"James", "Bond", "Shaken, not stirred"}
	moneypenny := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	slice := [][]string{james, moneypenny}

	for _, str := range slice {
		//fmt.Println(str)
		for _, words := range str {
			fmt.Printf("%s ", words)
		}
		fmt.Println()
	}

}
