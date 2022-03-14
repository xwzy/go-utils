package fs

import "os"

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Mkdir(path string) error {
	return os.MkdirAll(path, PermModeDir)
}

func MkdirIfNotExists(path string) error {
	if DirExists(path) {
		return nil
	}
	return Mkdir(path)
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func MoveDir(src, dst string) error {
	return os.Rename(src, dst)
}

func CopyDir(src, dst string) error {

}