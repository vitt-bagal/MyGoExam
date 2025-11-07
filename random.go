package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var base62 = []byte("012345689abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	shortId := make([]byte, 8)
	for i := range shortId {
		shortId[i] = base62[rand.Intn(len(base62))]
	}
	fmt.Println(string(shortId))
}
