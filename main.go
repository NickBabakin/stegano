package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/NickBabakin/stegano/filetools"
	"github.com/NickBabakin/stegano/standartstegano"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var bytesToShow = 160

func showHex(data []byte) {
	length := len(data)
	if length > bytesToShow {
		length = bytesToShow
	}
	j := 0
	for i := 0; i < length; i++ {
		if j == 16 {
			j = 0
			fmt.Print("\n")
		}
		fmt.Printf("%2.2X ", data[i])
		j++
	}
	fmt.Print("\n")
}

func main() {
	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	sourceFile := encryptCmd.String("f", "", "REQUIRED: source file used as a container")
	targetFile := encryptCmd.String("r", "result.bmp", "name of file to be created")
	msg := encryptCmd.String("m", "", "REQUIRED: message you want to encrypt")
	show := encryptCmd.Bool("s", false, "shows first 160 bytes of encypted and decrypted data")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptFile := decryptCmd.String("f", "", "REQUIRED: file to decrypt info from")

	if len(os.Args) < 2 {
		fmt.Print("expected 'encrypt' or 'decrypt' subcommands\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		if *sourceFile == "" || *targetFile == "" || *msg == "" {
			fmt.Print("f, m flags must be specified. See --help\n")
			os.Exit(1)
		}
		fileData, infoOffset, err := filetools.ParseBmp(*sourceFile)
		check(err)
		if *show {
			fmt.Print("Container before encryption\n")
			showHex(fileData)
		}
		err = standartstegano.PerformStandartEncryption(fileData[infoOffset:], []byte(*msg))
		check(err)
		filetools.WriteBmp(*targetFile, fileData)
		if *show {
			fmt.Print("Container after encryption\n")
			showHex(fileData)
		}
		fmt.Print("Message '", *msg, "' is encrypted.\nNew file name is ", *targetFile, "\n")

	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		if *decryptFile == "" {
			fmt.Print("f flag must be specified. See --help\n")
			os.Exit(1)
		}
		fileData, infoOffset, err := filetools.ParseBmp(*decryptFile)
		check(err)
		message, err := standartstegano.PerformStandartDecryption(fileData[infoOffset:])
		check(err)
		fmt.Print("Decryption ended successfully.\nThe message is '", string(message), "'\n")

	default:
		fmt.Print("Expected 'encrypt' or 'decrypt' subcommands\n")
		os.Exit(1)
	}
}
