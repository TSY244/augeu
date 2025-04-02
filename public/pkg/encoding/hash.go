package encoding

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Hash(str string) string {
	h := md5.New()
	md5Str := h.Sum([]byte(str))
	return hex.EncodeToString(md5Str)
}
