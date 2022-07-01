package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
)

func GetMd5Str(str string) string {
	md5 := md5.New()
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(str)
	md5.Write(buf.Bytes())
	return hex.EncodeToString(md5.Sum(nil))
}
