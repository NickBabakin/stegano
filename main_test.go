package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/NickBabakin/stegano/filetools"
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
	err := standartstegano.PerformStandartEncryption(container, message)
	testErr(err, t)
	fmt.Printf("Container with hidden message: \n - %s\n", container)

	fmt.Printf("\n--------------------------------\n")

	decryptedMessage, err := standartstegano.PerformStandartDecryption(container)
	testErr(err, t)
	fmt.Printf("Extracted from container message: \n - %s\n\n", decryptedMessage)
	if string(decryptedMessage) != string(message) {
		t.Errorf("\nexpected: %s\ngot: %s", decryptedMessage, message)
	}
}

func TestFullBmp(t *testing.T) {
	fileData, infoOffset, err := filetools.ParseBmp("abc.bmp")
	container := fileData[infoOffset:]
	testErr(err, t)

	message := []byte("Love")
	fmt.Printf("\nMessage: \n - %s\n", message)
	fmt.Printf("Container: \n - %s\n", container[:100])
	err = standartstegano.PerformStandartEncryption(container, message)
	testErr(err, t)
	fmt.Printf("Container with hidden message: \n - %s\n", container[:100])
	err = filetools.WriteBmp("file.bmp", fileData)
	testErr(err, t)

	fmt.Printf("\n--------------------------------\n")

	fileDataDec, infoOffset, err := filetools.ParseBmp("file.bmp")
	container = fileDataDec[infoOffset:]
	testErr(err, t)
	decryptedMessage, err := standartstegano.PerformStandartDecryption(container)
	testErr(err, t)
	fmt.Printf("Extracted from container message: \n - %s\n\n", decryptedMessage)
	if string(decryptedMessage) != string(message) {
		t.Errorf("\nexpected: %s\ngot: %s", decryptedMessage, message)
	}

	os.Remove("file.bmp")

}
