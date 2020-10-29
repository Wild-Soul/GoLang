package main

import "fmt"

const (
	_, current = iota, 2020
	y2021      = iota + current
	y2022      = iota + current
	y2023      = iota + current
	y2024      = iota + current
)

func main() {
	fmt.Println(y2021, y2022, y2023, y2024)
}
