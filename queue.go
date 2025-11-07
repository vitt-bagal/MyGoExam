package main

import "fmt"

func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	var hobbyists []string
	for ind, val := range hobbies {
		for _, v := range val {
			if v == hobby {
				hobbyists = append(hobbyists, ind)
			} else {
				continue
			}
		}
	}
	return hobbyists
}

func main() {
	hobbies := map[string][]string{
		"Steve": []string{"Fashion", "Piano", "Reading"},
		"Patty": []string{"Drama", "Magic", "Pets"},
		"Chad":  []string{"Puzzles", "Pets", "Yoga"},
	}

	fmt.Println(findAllHobbyists("Yoga", hobbies))
}
