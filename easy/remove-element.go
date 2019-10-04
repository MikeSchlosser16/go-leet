// Iterate through list
// if element is equal to the value
//        swap with the first non value
// else
//        keep the value, continue
func removeElement(nums []int, val int) int {

	for i, num := range nums {
		if num == val {
			// Find first non value
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != val {
					// Swap
					temp := nums[i]
					nums[i] = nums[j]
					nums[j] = temp
					break
				}
			}
		}
	}

	var good int
	for good = 0; good < len(nums); good++ {
		if nums[good] == val {
			break
		}
	}
	return good
}

// similar to above but much more simple
func removeElement(nums []int, val int) int {
    i := 0 // good elements, starting at index 0
    for j := 0; j < len(nums); j++{
        if nums[j] != val { // we have a number to keep
            nums[i] = nums[j] // keep it, adding to end of the list of good values at the front of array
            i += 1 // increment counter
        }
    }
    return i // i is the length of the new good array
}
