package filetools

import (
	"fmt"
	"log"
	"testing"
)

func TestFileToSlice(t *testing.T) {
	dataFromFile, err := FileToSlice("../abc.bmp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataFromFile[:100])
}
