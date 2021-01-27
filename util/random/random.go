package random

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	str := []byte("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	res := make([]byte, n)

	for i := 0; i < n; i++ {
		res[i] = str[rand.Intn(len(str))]
	}

	return string(res)
}
