package standartstegano

import (
	"encoding/binary"
	"errors"
)

func standartExtract(container []byte) ([]byte, error) {
	extractedData := make([]byte, len(container)/8)
	bitNumber := 0
	byteNumber := 0
	for containerCount := 0; containerCount < len(container); containerCount++ {
		bit := GetBit(container[containerCount], 7)
		extractedData[byteNumber] = extractedData[byteNumber] | bit<<(7-byte(bitNumber))
		if bitNumber == 7 {
			bitNumber = 0
			byteNumber++
		} else {
			bitNumber++
		}
	}
	return extractedData, nil
}

func PerformStandartExtraction(container []byte) ([]byte, error) {
	if len(container) < bytesOfSize*8 {
		return nil, errors.New("file corrupted")
	}
	sizeData, err := standartExtract(container[:bytesOfSize*8])
	if err != nil {
		return nil, err
	}
	sizeDataHelper := make([]byte, 4-len(sizeData), 4)
	sizeData = append(sizeDataHelper, sizeData...)
	dataLen := int(binary.BigEndian.Uint32(sizeData))
	if len(container) < bytesOfSize*8+dataLen*8 {
		return nil, errors.New("file corrupted")
	}
	extractedData, err := standartExtract(container[bytesOfSize*8 : bytesOfSize*8+dataLen*8])
	return extractedData, err
}
