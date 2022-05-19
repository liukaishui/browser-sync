package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func Md5(str string) (string, error) {
	hash := md5.New()
	if _, err := io.WriteString(hash, str); err != nil {
		return "", err
	}
	return string(hash.Sum(nil)), nil
}

func Md5File(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	dst := hex.EncodeToString(hash.Sum(nil))
	return dst, nil
}
