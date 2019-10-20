package main

import (
	"fmt"
)


// A few things to note here:
// 1. A common approach in string manipulation is to edit the string starting from the end and working backwards, this is useful because the extra buffer is at the end so we can change without fear of overwriting
// 2. Two scan approaches seem to be very common in string problems

func main(){
	fmt.Println(urlify([]rune("Mr John Smith    "), 13)) // 4 extra spaces at the end to fill
}

func urlify(s []rune, trueLength int) string {
	// First scan: count number of spaces. Get how many chars in the ending string
	var spaceCount, index int
	for i := 0; i < trueLength; i++{
		if string(s[i]) == " "{
			spaceCount++
		}
	}

	// Second scan: done in reverse order, edit the string. If space, replace, else keep
	index = trueLength + spaceCount * 2
	for j := trueLength - 1; j >= 0; j--{
		if string(s[j]) == " "{
			s[index - 1] = '0'
			s[index - 2] = '2'
			s[index - 3] = 'S'
			index -= 3
		} else{
			s[index - 1] = s[index]
			index--
		}
	}
	return string(s)
}