package util

import (
	"os"
	"testing"
)

func TestFileRead(t *testing.T) {
	dir, _ := os.Getwd()
	filename := dir + "/example.txt"

	got, err := FileRead(filename)
	if err != nil {
		t.Fatalf("contents cannot be got.")
	}

	want := "example"

	if got != want {
		t.Fatalf("want is %s, but got is %s", want, got)
	}
}
