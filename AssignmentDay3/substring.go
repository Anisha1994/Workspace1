package main

import "fmt"

func main() {
	value := "AssignmentDay3"
	runes := []rune(value)
	for length := 1; length <= len(runes); length++ {
		for start := 0; start <= len(runes) - length; start++ {
			substring := string(runes[start:start+length])
			fmt.Println(substring)
		}
	}
}