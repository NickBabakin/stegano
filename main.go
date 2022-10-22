package main

import (
	"fmt"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello World!")
}
