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

func FileWrite(filename string, data []byte) error {
	var f *os.File
	var err error
	if exists(filename) {
		f, err = os.Open(filename)
	} else {
		f, err = os.Create(filename)
	}

	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
