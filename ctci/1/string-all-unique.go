package main

import (
	"fmt"
)

func main(){
	fmt.Println(allUnique("hello"))
	fmt.Println(allUnique("michael"))
}

// Time: O(n) -- or O(1)?
// Space: O(n)
func allUnique(s string) bool {
	chMap := make(map[string]bool)
	for _, ch := range s{
		if _, found := chMap[string(ch)]; found {
			return false
		}else{
			chMap[string(ch)] = true
		}
	}
	return true
}