package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// stringWithCharset 랜덤 문자열 생성
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString 랜덤 문자열 생성
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	return stringWithCharset(rand.Intn(length)+1, charset)
}
