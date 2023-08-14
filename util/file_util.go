package util

import (
	"io"
	"os"
)

func FileRead(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
