package handler

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
)

// n is the length of random string we want to generate
func randStr(n int) string {
	var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func generatePasswordHash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func validatelogin(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}