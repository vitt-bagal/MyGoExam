package main

import "fmt"

func main() {
	str1 := "a,b$z"

	for i, c := range str1 {
		fmt.Println(i, " => ", string(c))
	}
	for i, r := range str1 {
		fmt.Printf("Type of i is %T and r is %T \n", i, r)
		fmt.Println(i, r, string(r))
	}
	fmt.Printf("Reverse of string %v is %v \n", str1, revStringchar(str1))
	//str2 := "Ab,c,de!$"
	//fmt.Printf("Reverse of string %v is %v \n", str2, revStringchar(str2))
}

// https://www.geeksforgeeks.org/reverse-a-string-without-affecting-special-characters/
//function to reverse string witout affecting special character
func revStringchar(str string) string {
	var b byte
	l, r := 0, len(str)-1
	// Loop over the string with len.
	for l < r {
		if !isAlpha(str[l]) {
			l++
		} else if !isAlpha(str[r]) {
			r--
		} else {
			fmt.Printf("Type of str[l] is %T", str[l])
			b = str[l]
			str[l] = str[r]
			str[r] = b
		}
	}
	return str
}

func isAlpha(r byte) bool {
	if r >= 65 && r <= 90 || r >= 97 && r <= 122 {
		return true
	} else {
		return false
	}
}
