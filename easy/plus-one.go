// Two cases - 9 or non 9
// Non 9 is not interesting
// On 9, we need to make digit 0 and it will then increment the NEXT left element in the array
// Don't overthink the 9 case..
func plusOne(digits []int) []int {
    for i := len(digits)-1; i >= 0; i--{
        if digits[i] < 9{
            digits[i] += 1
            return digits
        }else{
            digits[i] = 0
        }
    }
    // All digits must be 9 to get here
    val := make([]int, len(digits)+1)
    val[0] = 1
    return val
}
