package fs

import (
	"fmt"
	"testing"
)

func TestWriteFile(t *testing.T) {
	str := "Hello, world!"
	err := WriteFile("/tmp/test.txt", []byte(str))
	if err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	str := "Hello, world!"
	err := WriteFile("/tmp/test.txt", []byte(str))
	if err != nil {
		t.Error(err)
	}
	data, err := ReadFile("/tmp/test.txt")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(data))
	if string(data) != str {
		t.Error("ReadFile failed")
	}
}
