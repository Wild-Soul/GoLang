package main

import (
	"fmt"
)

type customErr struct {
	msg string
}

func (err customErr) Error() string {
	return err.msg
}

func main() {
	c := customErr{
		msg: "Dummy error",
	}

	foo(c)
}

func foo(err error) {
	fmt.Println(err)
}
