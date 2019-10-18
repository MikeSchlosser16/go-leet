package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	fmt.Println(isPermutation("God", "dog"))
	fmt.Println(isPermutation("god", "dog"))
}


// Strings of different lengths can not be permutations of each other..
// Thus, we can sort and then compare sorted lists
// We could also store values in a map of [char]count and decrement count as we check through, if < 0 return false
func isPermutation(val1, val2 string) bool {
	val1Chars := strings.Split(val1, "")
	sort.Strings(val1Chars)

	val2Chars := strings.Split(val2, "")
	sort.Strings(val2Chars)

	if len(val1Chars) != len(val2Chars){
		return false
	}

	for i := 0; i < len(val1Chars); i++{
		if val1Chars[i] != val2Chars[i]{
			return false
		}
	}
	return true
}