package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func Md5Encrypt(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashFile(file *os.File) (hashStr string, err error) {
	defer file.Close()
	hash := sha256.New()
	if _, err = io.Copy(hash, file); err != nil{
		return
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

func HashByte(b []byte) string {
	hash := sha1.New()
	hash.Write(b)
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
