package encoding

import "crypto/md5"

func Md5Hash(str string) string {
	md5Str := md5.Sum([]byte(str))
	return string(md5Str[:])
}
