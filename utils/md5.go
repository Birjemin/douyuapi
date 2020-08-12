package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
)

// GetMD5String md5 string
func GetMD5String(strings string) string {

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strings))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// Md5ByByte md5 by byte
func Md5ByByte(bytes []byte) string {

	md5Ctx := md5.New()
	md5Ctx.Write(bytes)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// Bytes generates n random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// Hex ...
func Hex(n int) string {
	return hex.EncodeToString(Bytes(n))
}
