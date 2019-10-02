func longestCommonPrefix(strs []string) string {

	// Handle edge case
	if len(strs) == 0 {
		return ""
	}

	var prefix string

	for i, ch := range strs[0] {

		for _, word := range strs[1:] {
			// Handle case where length of element is shorter than the index we are on, this is upper bound for the prefix
      // If the char does not match, we also know this is the limit of the prefix
			if len(word) < i+1 || rune(word[i]) != ch {
				return prefix
			}
		}
    // If we made it here, we were valid through all words so append that rune to the prefix
		prefix += string(ch)
	}
	return prefix
}

// Could also be done with binary search
// Note: longest possible common prefix is the len of the shortest word
