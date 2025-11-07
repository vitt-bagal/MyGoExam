package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `mumbai is capital city of maharashtra and mumbai is big city in india`
	m := make(map[string]int)
	s := strings.Split(str, " ")
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if ok {
			m[s[i]] = v + 1
		} else {
			m[s[i]] = 1
		}
	}
	fmt.Println("Map- ", m)
}
