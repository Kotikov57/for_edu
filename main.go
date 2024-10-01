package main

import (
	"fmt"
)

func main() {
	var j, s string
	fmt.Scan(&j)
	fmt.Scan(&s)
	count := 0
    seen := make(map[rune]bool)
    result := ""
    for _, char := range j {
        if !seen[char] {
            result += string(char)
            seen[char] = true
        }
    }
	for i := range result{
		for k := range s{
			if result[i] == s[k]{
				count++
			}
		}
	}
	fmt.Println(count)
}
