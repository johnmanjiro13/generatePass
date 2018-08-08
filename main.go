package main

import (
	"fmt"
	"math/rand"
	"time"
)

const useLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = useLetters[rand.Intn(len(useLetters))]
	}
	return string(b)
}

func main() {
	fmt.Println(RandString(12))
}
