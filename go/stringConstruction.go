package main

import "fmt"

// stringConstruction returns how many times can we constuct the string a from the string b
func stringConstruction(a, b string) int {
	count := 0

	repetitions := make(map[rune]int)
	for _, char := range b {
		repetitions[char]++
	}

	for {
		for _, char := range a {
			if repetitions[char] > 0 {
				repetitions[char]--
			} else {
				return count
			}
		}
		count++
	}
}

func main() {
	var a, b string
	fmt.Scanf("%s\n", &a)
	fmt.Scanf("%s\n", &b)
	fmt.Println(stringConstruction(a, b))
}
