// Program to validate brackets
// {A+B} -- valid
// [A+{B+C} --> Invalid // square bracket do not have closing bracket.
// }{ --> Invalid // shd not star with close
package main

import (
	"fmt"
	"strings"
)

func main() {
	//listStrings := []string{"[A+(B+C)", "{A+B}", "}{"}
	str := "}{"
	checkValid(str)
	//listStrings := []string{"[A+(B+C)", "{A+B}", "}{"}

}

func checkValid(s string) {
	count := 0
	for i := 0; i < len(s); i++ {
		if strings.Contains("{[(", string(s[i])) {
			count++
		} else if strings.Contains("]})", string(s[i])) {
			count--
			if i == 0 {
				break
			}
		} else {
			continue
		}
	}

	if count == 0 {
		fmt.Println("valid...")
	} else {
		fmt.Println("not valid...")
	}
}
