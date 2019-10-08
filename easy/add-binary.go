func addBinary(a string, b string) string {
	aI := binaryToInt(a)
	bI := binaryToInt(b)
	total := aI + bI
	return intToBinary(total)

}

func binaryToInt(binary string) int {
	var val, exponentCounter int
	for i := len(binary) - 1; i >= 0; i-- {
		val += int(binary[i]-'0') * int(math.Pow(2, float64(exponentCounter)))
		exponentCounter++
	}
	return val
}

func intToBinary(val int) string {
	//var vals []string
	//for val != 0 {
	//	remainder := val % 2
	//	vals = append([]string{string(remainder)}, vals...)
	//	val /= 2

	//}
	//var strVal string
	//for _, ch := range vals {
	//	strVal += string(ch)
	//}
	//return strVal
	v := strconv.FormatInt(int64(val), 2) // 1111011
	return v

}




// Java implementation
class Solution {
    public String addBinary(String a, String b) {
        StringBuilder sb = new StringBuilder();
        int i = a.length() - 1;
        int j = b.length() - 1;
        int carry = 0;
        while (i >= 0 || j >= 0){
            int sum = carry;
            if (i >= 0) {
                sum += a.charAt(i--) - '0'; // convert to int
            }
            if (j >= 0){
                sum += b.charAt(j--) - '0';
            }

            sb.insert(0, sum % 2); // sum can be 0,1,2
            carry = sum / 2; // only get a val if sum is 2, else int division to 0
        }
        // Check if last sig digits need to account for carry
        if (carry > 0){
            sb.insert(0,1); // add the final one
        }
        return sb.toString();
    }
}
