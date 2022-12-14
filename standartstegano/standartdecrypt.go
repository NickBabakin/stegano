package standartstegano

import (
	"encoding/binary"
	"errors"
)

func standartDecrypt(container []byte) ([]byte, error) {
	decryptedData := make([]byte, len(container)/8)
	bitNumber := 0
	byteNumber := 0
	for containerCount := 0; containerCount < len(container); containerCount++ {
		bit := GetBit(container[containerCount], 7)
		decryptedData[byteNumber] = decryptedData[byteNumber] | bit<<(7-byte(bitNumber))
		if bitNumber == 7 {
			bitNumber = 0
			byteNumber++
		} else {
			bitNumber++
		}
	}
	return decryptedData, nil
}

func PerformStandartDecryption(container []byte) ([]byte, error) {
	if len(container) < bytesOfSize*8 {
		return nil, errors.New("file corrupted")
	}
	sizeData, err := standartDecrypt(container[:bytesOfSize*8])
	if err != nil {
		return nil, err
	}
	sizeDataHelper := make([]byte, 4-len(sizeData), 4)
	sizeData = append(sizeDataHelper, sizeData...)
	dataLen := int(binary.BigEndian.Uint32(sizeData))
	if len(container) < bytesOfSize*8+dataLen*8 {
		return nil, errors.New("file corrupted")
	}
	decryptedData, err := standartDecrypt(container[bytesOfSize*8 : bytesOfSize*8+dataLen*8])
	return decryptedData, err
}
