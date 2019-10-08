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
