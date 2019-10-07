
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

// Java
class Solution {
    public String countAndSay(int n) {
        // validation
        if (n <= 0){
            return "";
        }
        String res = "1"; // Default return value to 1, must always start here
        while(n > 1){ // Start at n, go back to one, which we know
            StringBuilder cur = new StringBuilder(); // Cur holds the value of the current term
            for (int i = 0; i < res.length(); ++i){  // Loop through the current steps result
                int count = 1;
                while(i + 1 < res.length() && res.charAt(i) == res.charAt(i+1)){ // Ensure we arent out of bouns and that this element is the same as the next. If they are the same, add to counter and increment i again
                    ++count;
                    ++i;
                }
                cur.append(count).append(res.charAt(i));
            }
            res = cur.toString();
            --n;
        }
        return res;
    }
}
