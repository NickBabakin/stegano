package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/NickBabakin/stegano/standartstegano"
)

func testErr(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestFullWithText(t *testing.T) {
	container := []byte("Hello, World! This is my first day here. I am so glad to see all of you")
	message := []byte("Love")
	fmt.Printf("\nMessage: \n - %s\n", message)
	fmt.Printf("Container: \n - %s\n", container)
	bytesOfSize := 2
	err := standartstegano.PerformStandartEncryption(container, message, bytesOfSize)
	testErr(err, t)
	fmt.Printf("Container with hidden message: \n - %s\n", container)

	fmt.Printf("\n--------------------------------\n")

	decryptedMessage, err := standartstegano.PerformStandartDecryption(container, bytesOfSize)
	testErr(err, t)
	fmt.Printf("Extracted from container message: \n - %s\n\n", decryptedMessage)
	if string(decryptedMessage) != string(message) {
		t.Errorf("\nexpected: %s\ngot: %s", decryptedMessage, message)
	}
}

func TestFullBmp(t *testing.T) {
	fileData, err := ioutil.ReadFile("abc.bmp")
	container := fileData[100:]
	testErr(err, t)

	message := []byte("Love")
	fmt.Printf("\nMessage: \n - %s\n", message)
	fmt.Printf("Container: \n - %s\n", container[:100])
	bytesOfSize := 2
	err = standartstegano.PerformStandartEncryption(container, message, bytesOfSize)
	testErr(err, t)
	fmt.Printf("Container with hidden message: \n - %s\n", container[:100])
	err = ioutil.WriteFile("file.bmp", fileData, 0600)
	testErr(err, t)

	fmt.Printf("\n--------------------------------\n")

	fileDataDec, err := ioutil.ReadFile("file.bmp")
	container = fileDataDec[100:]
	testErr(err, t)
	decryptedMessage, err := standartstegano.PerformStandartDecryption(container, bytesOfSize)
	testErr(err, t)
	fmt.Printf("Extracted from container message: \n - %s\n\n", decryptedMessage)
	if string(decryptedMessage) != string(message) {
		t.Errorf("\nexpected: %s\ngot: %s", decryptedMessage, message)
	}

	os.Remove("file.bmp")

}
