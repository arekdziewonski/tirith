package totp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHmacSha1(t *testing.T) {
	secret := "12345678901234567890"

	assert.Equal(t, "cc93cf18508d94934c64b65d8ba7667fb7cde4b0", hmacSha1(secret, 0))
	assert.Equal(t, "75a48a19d4cbe100644e8ac1397eea747a2d33ab", hmacSha1(secret, 1))
	assert.Equal(t, "0bacb7fa082fef30782211938bc1c5e70416ff44", hmacSha1(secret, 2))
	assert.Equal(t, "66c28227d03a2d5529262ff016a1e6ef76557ece", hmacSha1(secret, 3))
	assert.Equal(t, "a904c900a64b35909874b33e61c5938a8e15ed1c", hmacSha1(secret, 4))
	assert.Equal(t, "a37e783d7b7233c083d4f62926c7a25f238d0316", hmacSha1(secret, 5))
	assert.Equal(t, "bc9cd28561042c83f219324d3c607256c03272ae", hmacSha1(secret, 6))
	assert.Equal(t, "a4fb960c0bc06e1eabb804e5b397cdc4b45596fa", hmacSha1(secret, 7))
	assert.Equal(t, "1b3c89f65e6c9e883012052823443f048b4332db", hmacSha1(secret, 8))
	assert.Equal(t, "1637409809a679dc698207310c8c7fc07290d9e5", hmacSha1(secret, 9))
}
