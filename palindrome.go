package main

import "fmt"

func main() {
	fmt.Println(IsPalindrome("tamat"))
}

func IsPalindrome(input string) bool {
	var inputLenDiv2 = len(input) / 2
	var inputLenFull = len(input)

	for i := 0; i <= inputLenDiv2; i++ {
		if input[i] == input[inputLenFull-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}
