// This takes extra space, since we need to save the string
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    stringVal := fmt.Sprintf("%v", x) // can't just convert, consider zeros
    for i := 0; i < len(stringVal); i++{
        if stringVal[i] != stringVal[len(stringVal)-(i+1)]{
            return false
        }
    }
    return true
}

// Time complexity: O(log(n)). We devided input by 10 every iteration
// Space complexity: O(1)
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    // To be a palindrome, if last digit is 0 to be a palindrome is must be zero(nothing else satifies this case)
    if x % 10 == 0 && x != 0 {
        return false
    }

    /* Example:
        x = 1221    reversed = 0
        x = 122     reversed = (0 * 10) + 1 = 1
        x = 12      reversed = (1 * 10) + 2 = 12
    */

    // Get last digit by mod 10, then remove the last digit by diving by 10
    var reversed int
    for(x > reversed){ // once we are half way, we have reversed the entire thing
        reversed = (reversed * 10) + (x % 10)
        x /= 10
    }

    return x == reversed || x == reversed / 10 // Divide by 10 handles oddd number, ie 12321. x here = 12, reverted = 123 but middle digit doesnt matter in palindrome

}
