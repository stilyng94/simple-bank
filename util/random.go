package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@mail.com", RandomString(6))
}

func RandomPassword() string {
	return fmt.Sprintf("test.%s", RandomString(5))
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, GHS, POUND, CAD}
	return currencies[rand.Intn(len(currencies))]
}

func RandomAmount() float64 {
	return float64(rand.Int63n(100000))
}
