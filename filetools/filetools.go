package filetools

import (
	"encoding/binary"
	"io/ioutil"
)

func ParseBmp(filename string) ([]byte, int, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, 0, err
	}
	infoOffset := int(binary.LittleEndian.Uint32(fileData[10:14]))
	return fileData, infoOffset, nil
}

func WriteBmp(filename string, fileData []byte) error {
	if err := ioutil.WriteFile(filename, fileData, 0600); err != nil {
		return err
	}
	return nil
}
