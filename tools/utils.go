package tools

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

func GetMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return strings.ToLower(hex.EncodeToString(m.Sum(nil)))
}

func GetUUID() string {
	u := uuid.NewV4()
	return u.String()
}
func RandomStr(lenNum int) string {
	var chars = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	str := strings.Builder{}
	length := len(chars)
	rand.Seed(time.Now().UnixNano()) //重新播种，否则值不会变
	for i := 0; i < lenNum; i++ {
		str.WriteString(chars[rand.Intn(length)])

	}
	return str.String()
}

// DeepCopyStr 用于深度复制str
func DeepCopyStr(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func ifs[T any](a bool, b, c T) T {
	if a {
		return b
	}
	return c
}
