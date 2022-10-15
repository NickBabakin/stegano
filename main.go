package main

import (
	"fmt"
	"log"

	"github.com/NickBabakin/stegano/standartstegano"
)

func main() {
	container := []byte("Hello, World! This is my first day here. I am so glad to see all of you")
	data := []byte("Love")
	fmt.Printf("\nMessage: \n - %s\n", data)
	fmt.Printf("Container: \n - %s\n", container)
	bytesOfSize := 2
	err := standartstegano.PerformStandartEncryption(container, data, bytesOfSize)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Container with hidden message: \n - %s\n", container)
	fmt.Printf("\n--------------------------------\n")
	decryptedInfo, err := standartstegano.PerformStandartDecryption(container, bytesOfSize)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Extracted from container message: \n - %s\n\n", decryptedInfo)
}
