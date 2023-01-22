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
	hideCmd := flag.NewFlagSet("hide", flag.ExitOnError)
	sourceFile := hideCmd.String("f", "", "REQUIRED: source file used as a container")
	targetFile := hideCmd.String("r", "result.bmp", "name of file to be created")
	msg := hideCmd.String("m", "", "REQUIRED: message you want to hide")
	show := hideCmd.Bool("s", false, "shows first 160 bytes of source container and container with hidden data")

	extractCmd := flag.NewFlagSet("extract", flag.ExitOnError)
	extractFile := extractCmd.String("f", "", "REQUIRED: file to extract info from")

	if len(os.Args) < 2 {
		fmt.Print("expected 'hide' or 'extract' subcommands\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hide":
		hideCmd.Parse(os.Args[2:])
		if *sourceFile == "" || *targetFile == "" || *msg == "" {
			fmt.Print("f, m flags must be specified. See --help\n")
			os.Exit(1)
		}
		fileData, infoOffset, err := filetools.ParseBmp(*sourceFile)
		check(err)
		if *show {
			fmt.Print("Container before hiding\n")
			showHex(fileData)
		}
		err = standartstegano.PerformStandartHiding(fileData[infoOffset:], []byte(*msg))
		check(err)
		filetools.WriteBmp(*targetFile, fileData)
		if *show {
			fmt.Print("Container after hiding\n")
			showHex(fileData)
		}
		fmt.Print("Message '", *msg, "' is hidden.\nNew file name is ", *targetFile, "\n")

	case "extract":
		extractCmd.Parse(os.Args[2:])
		if *extractFile == "" {
			fmt.Print("f flag must be specified. See --help\n")
			os.Exit(1)
		}
		fileData, infoOffset, err := filetools.ParseBmp(*extractFile)
		check(err)
		message, err := standartstegano.PerformStandartExtraction(fileData[infoOffset:])
		check(err)
		fmt.Print("Extraction ended successfully.\nThe message is '", string(message), "'\n")

	default:
		fmt.Print("Expected 'hide' or 'extract' subcommands\n")
		os.Exit(1)
	}
}
