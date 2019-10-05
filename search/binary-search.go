// classic divide and conquer
func binarySearch(vals []int, target, left, right int) (index int, found bool) {
	mid := int((left + right) / 2)
	if vals[mid] == target {
		return mid, true
	}

	if vals[mid] < target {
		// search right
		return binarySearch(vals, target, mid+1, right)
	} else {
		// search left
		return binarySearch(vals, target, left, mid-1)
	}

}
