package main

import "fmt"

func main() {
	str := "Geeks"
	reverseWord(str)
	//fmt.Println("Reverse string is", reverseWord(str))
}

func reverseWord(s string) {
	rev := make([]byte, len(s))
	for i, j := 0, len(s)-1; j >= 0; i, j = i+1, j-1 {
		rev[i] = s[j]
	}
	fmt.Println((string)(rev))
	//buf := []rune(s)
	//for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
	//	buf[i], buf[j] = buf[j], buf[i]
	//}
	//fmt.Println((string)(buf))
}
