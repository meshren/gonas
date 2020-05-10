package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Encrypt(str string) string {
	h := md5.New()
	_, _ = io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
