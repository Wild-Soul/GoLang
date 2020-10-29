package main

import "fmt"

func main() {

	details := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	//fmt.Println(details)
	fmt.Printf("%T\n", details)

	for k, v := range details {
		fmt.Println("Details for ", k)
		for _, s := range v {
			fmt.Printf("\t%s", s)
		}
		fmt.Println()
	}

	fmt.Println("===========================delete called==============================")
	delete(details, "bond_james")

	for k, v := range details {
		fmt.Println("Details for ", k)
		for _, s := range v {
			fmt.Printf("\t%s", s)
		}
		fmt.Println()
	}
}
