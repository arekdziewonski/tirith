package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func hmacSha1(secret string, count int) string {
	text := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		text[i] = (byte)(count & 0xff)
		count >>= 8
	}

	key := []byte(secret)
	h := hmac.New(sha1.New, key)
	h.Write(text)
	s := h.Sum(nil)
	return fmt.Sprintf("%x", s)
}
