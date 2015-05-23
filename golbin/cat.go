package golbin

import (
	"io/ioutil"
)

/*
Cat returns full content of a file as string.
*/
func Cat(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
