package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Black",
		},
		fourWheels: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
}
