func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	lower := strings.ToLower(s)

	cleanedStr := reg.ReplaceAllString(lower, "")


	for i := 0; i < len(cleanedStr) / 2; i++{
		if cleanedStr[i] != cleanedStr[len(cleanedStr) - i - 1]{
			return false
		}
	}
	return true
}