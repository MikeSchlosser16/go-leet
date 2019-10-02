// The bad hacky convert to string method - use math, not string!
func reverse(x int) int {
    // convert int to string to reverse
    strVal := strconv.Itoa(x)
    var reverseStr string
    for i := len(strVal)-1; i >= 0; i--{
        reverseStr += string(strVal[i])
    }

    // Handle negative case - could also use a bool and multiply by 1
    if string(reverseStr[len(reverseStr)-1]) == "-"{
        reverseStr = "-" + reverseStr[0:len(reverseStr)-1]
    }

    // Reverse string
    val, err := strconv.Atoi(reverseStr)
    if err != nil {
        return 0
    }

    // Handle over flow from prompt
    if val > math.MaxInt32 || val < math.MinInt32{
        return 0
    }
    return val
}

// Reverse using math. Pop the last digit off the stack using mod 10, ie 123 % 10 = 3. Then divide by 10, 12 % 10 = 2, 1 % 10 = 1
// The mod also takes care of the negative, ie. -321 % 10 == -1
// Time complexity: O(log(x))
// Space complexity: O(1)

// x = 123
// mod = 123 % 10 = 3
// x = 12
// reverse = 0 * 10 + 3 = 3
// mod = 12 % 10 = 2
// x = 1
// reverse = 3 * 10 + 2 = 32
// mod = 1 % 10 = 1
// x = 0
// reverse = 32 * 10 + 1 = 321

func reverse(x int) int {
    reverse := 0
    for x != 0 {
        mod := x % 10
        x /= 10
        // Check for overflow
        reverse = reverse * 10 + mod
        if reverse > math.MaxInt32  || reverse < math.MinInt32{
            return 0
        }
    }
    return reverse
}
