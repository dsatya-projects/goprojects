package main

import (
	"fmt"

	"github.com/satya-dillikar/goprojects/module1"
)

func main() {
	// Get a greeting message and print it.
	message, err := module1.Hello("Satya")
	if err == nil {
		fmt.Println(message)
	}
	fmt.Println("API return value", module1.BestMascot())
}
