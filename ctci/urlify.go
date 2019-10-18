package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(urlify("hello  there"))
	fmt.Println(urlify("michael"))
}

// Time: O(n) -- or O(1)?
// Space: O(n)
func urlify(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}