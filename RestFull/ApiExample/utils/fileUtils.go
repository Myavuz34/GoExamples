package utils

import (
	"io/ioutil"
)

//LoadFile Dosya okuma metotu
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
