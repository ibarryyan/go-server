package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func GetMd5Str(str string) string {
	md5 := md5.New()
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(str)
	md5.Write(buf.Bytes())
	return hex.EncodeToString(md5.Sum(nil))
}

func GetTokenStr() string {
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	c := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	var str string = ""
	for i := 1; i <= 18; i++ {
		str = str + charArr[ran.Intn(c)]
	}
	return str
}
