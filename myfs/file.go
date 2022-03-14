package myfs

import (
	"bytes"
	"io"
	"os"
)

func WriteFile(filename string, data []byte) error {
	openFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, PermModeFile)
	if err != nil {
		return err
	}
	defer openFile.Close()
	_, err = openFile.Write(data)
	return err
}

func ReadFile(filename string) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	openFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer openFile.Close()
	_, err = io.Copy(buf, openFile)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func RemoveFile(filename string) error {
	return os.Remove(filename)
}

func MoveFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	return err
}

