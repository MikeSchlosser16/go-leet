func removeDuplicates(nums []int) int {
	index := 1 // initial is always unique
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] != nums[i+1] {  // Next one isnt the same, need to add to the return value
			nums[index] = nums[i+1] // Set that value so we could upon return just take a selection of initial nums
			index += 1

		}
	}
	return index // index is now number of unique elements
}
