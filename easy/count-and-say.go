
// Use previous step, say COUNT:ELEMENT
func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    var val string
    temp := countAndSay(n-1)
    count := 1

    for i := 1; i < len(temp); i++{
        // last element
        last := string(temp[i-1])

        // We "get out" by either reaching the end of the string, or a new character

        // Already at the end, add the count and last element
        if i == len(temp){
            val += string(count) + last
        // Current element is different from the last, add the count and last element
        } else if string(temp[i]) != last{
            val += string(count) + last
            count = 1
        // Current element is the same, increment the count in COUNT:ELEMENT
        } else {
            count += 1
        }
    }
    return val
}
