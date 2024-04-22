package random

import (
	"math/rand"
	"time"
)

func NewRandomString(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	alphabet := []rune(lowercase + uppercase + numbers)

	res_runes := make([]rune, length)

	for ind := range res_runes {
		randInd := rnd.Intn(len(alphabet))
		res_runes[ind] = alphabet[randInd]
	}

	return string(res_runes)
}
