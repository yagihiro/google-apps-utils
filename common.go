package main

import (
	"math/rand"
	"time"
)

// RandomString は初期パスワード等で利用することを想定した length 長の文字列を返す関数です
func RandomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
