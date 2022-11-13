package filetools

import (
	"io/ioutil"
)

func FileToSlice(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}
