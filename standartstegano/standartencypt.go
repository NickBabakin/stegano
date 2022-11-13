package standartstegano

import (
	"encoding/binary"
	"errors"
	"math"
)

var bytesOfSize = 2

func GetBit(b byte, bitNumber int) byte {
	return (b >> (7 - bitNumber) & 1)
}

func powInt(x, y int) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

func checkCapacity(containerLen int, dataLen int, bytesOfSize int) error {
	if powInt(2, 8*bytesOfSize) < int64(dataLen) {
		return errors.New("not enough bits to encrypt size")
	}
	if containerLen < (dataLen+bytesOfSize)*8 {
		return errors.New("container too small to fit this data")
	}
	return nil
}

func standartEncrypt(container []byte, data []byte) error {
	containerCount := 0
	for byteNumber := 0; byteNumber < len(data); byteNumber++ {
		//fmt.Printf("%.8b\n", data[byteNumber])
		for bitNumber := 0; bitNumber < 8; bitNumber++ {
			bit := GetBit(data[byteNumber], bitNumber)
			if bit == 0 {
				container[containerCount] = container[containerCount] & 254
			} else {
				container[containerCount] = container[containerCount] | 1
			}
			containerCount++
		}
	}
	return nil
}

func PerformStandartEncryption(container []byte, data []byte) error {
	if err := checkCapacity(len(container), len(data), bytesOfSize); err != nil {
		return err
	}
	sizeData := make([]byte, 4)
	binary.BigEndian.PutUint32(sizeData, uint32(len(data)))
	if err := standartEncrypt(container[:bytesOfSize*8], sizeData[4-bytesOfSize:4]); err != nil {
		return err
	}
	if err := standartEncrypt(container[bytesOfSize*8:], data); err != nil {
		return err
	}
	return nil
}
