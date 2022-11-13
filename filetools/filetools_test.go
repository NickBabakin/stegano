package filetools

import (
	"fmt"
	"log"
	"testing"
)

func TestParseBmp(t *testing.T) {
	dataFromFile, infoOffset, err := ParseBmp("../abc.bmp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dataFromFile[:100])
	fmt.Println("offset: ", infoOffset)
	if infoOffset != 54 {
		t.Errorf("expected: 54\ngot: %d", infoOffset)
	}
}
