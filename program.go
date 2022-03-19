package main

// user 2 string
// st = abc
// str = bac

// acc or cbb
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "abc"
	str2 := "bac"
	if checkStringEqual(str1, str2) {
		fmt.Println("Both strings are equal")
	} else {
		fmt.Println("Both strings are not equal")
	}

}

func checkStringEqual(s1, s2 string) bool {
	if utf8.RuneCountInString(s1) == utf8.RuneCountInString(s2) {

	}
	return true
}
