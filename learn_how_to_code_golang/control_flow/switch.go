package main

import "fmt"

func main() {
	year := 2006

	switch year {
	case 2006:
		fmt.Println("Recovery")
	case 2000:
		fmt.Println("Marshall Mathers LP")
	case 2018:
		fmt.Println("Kamikaze")
	default:
		fmt.Println("Comon dude listen to his songs")
	}
}
