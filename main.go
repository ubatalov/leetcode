package main

import (
	"fmt"
	"leetcode/problems/easy"
)

func main() {
	fmt.Println(easy.IsValid("{{[()]}}"))
	fmt.Println(easy.IsValid("([)]"))
	fmt.Println(easy.IsValid("()[]{}"))
}
