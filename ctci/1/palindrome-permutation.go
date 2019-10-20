package main

import "fmt"

func main(){
	//fmt.Println(getPermutations("abc"))
	//fmt.Println(isPalindrome("civic"))
	//fmt.Println(isPalindrome("bad"))
	fmt.Println(isPermutationOfPalindrome("ivicc"))
	fmt.Println(isPermutationOfPalindromeGood("ivicc"))
}

/*
Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards.
A permutation is a rearrangement of letters.
*/

// Need to figure out what this problem is really asking..
// 'civic' -> true, its a palindrome
// 'ivicc' -> true, one of it's permutations, 'civic', is a palindrome
// 'civil' -> false, none of its permutations are palindromes
// 'livci' -> false, none of its permutations are palindromes

// The obvious approach is to compute all permutations of a string, then check each of those and see if that is a palindrome
//func permutationPalindrome(s string) bool {
//
//}

func isPermutationOfPalindrome(str string) bool{
	permutations := getPermutations(str)

	for _, permutation := range permutations{
		if isPalindrome(permutation){
			return true
		}
	}
	return false
}


// civic
// i = 0, str[0] = c, str[5-0-1] = c
// i = 1, str[1] = i, str[5-1-1] = i
// i = 2, str[2] = v, str[5-2-1] = v

func isPalindrome(str string) bool{
	for i := 0; i < len(str) / 2; i++{
		if str[i] != str[len(str) - i - 1]{
			return false
		}
	}
	return true
}



/*
 The permutation of 1 element is 1 element
 The permutation of a set of elements is a list of each of the elements concatenated with every permutation of the other elements.

 getPermutations("a") --> ["a"]
 getPermutations("ab") -->  "a" + perm(b)
 							"b" + perm(a)
  getPermutations("abc") --> "a" + perm(bc)
 							 "b" + perm(ac)
 							 "c" + perm(ab)
*/


func getPermutations(str string) []string{

	fmt.Println("in getPermutations for string: ", str)

	if len(str) == 1{
		fmt.Println("call to getPermutations returned: ", []string{str})
		return []string{str}
	}

	current := str[0:1]
	remainingString := str[1:]

	permutations := getPermutations(remainingString)

	allPermutations := make([]string, 0)

	for _, permutation := range permutations{
		for i := 0; i <= len(permutation); i++{
			// Insert char
			fmt.Println("inserting: ", current, " on string: ", permutation, " at index: ", i)
			newPermutation := insertAt(i, current, permutation)
			allPermutations = append(allPermutations, newPermutation)
		}
	}
	fmt.Println("call to getPermutations returned: ", allPermutations)
	return allPermutations
}

func insertAt(position int, charToInsert, permutation string) string{
	beginning := permutation[0:position]
	end := permutation[position:]
	return beginning + charToInsert + end
}




/////////////////////////////////////////////
/////////////////////////////////////////////
/////////////////////////////////////////////

// Better approach
// To determine if a string is a permutation of a palindrome, what does that mean?
// we need to know if it can be written the same forward and backward
// We need to have an even number of all characters or an odd number with 1 characters count being odd
// For example, we know tactcoapapa is a permutation
// [t: 2, a: 4, c: 2, o: 1, p: 2]
// All even except 1, o, so this is a permutation
// To be more precise, strings with even length (after removing all non-letter characters) must have all even counts of characters. Strings of an odd length must have exactly one character with an odd count.
// . It's therefore sufficient to say that, to be a permutation of a palindrome, a string can have no more than one character that is odd. This will cover both the odd and the even cases.
func isPermutationOfPalindromeGood(str string) bool{
	characterCounts := make(map[string]int)
	for _, ch := range str{
		if characterCounts[string(ch)] != 0{
			characterCounts[string(ch)]++
		}else{
			characterCounts[string(ch)] = 1
		}
	}

	var numOdds, numEvens int
	for _, count := range characterCounts{
		if count % 2 == 0{
		numEvens++
		}else {
		numOdds++
		}
	}

	if numOdds > 1{
		return false
	}
	return true
}