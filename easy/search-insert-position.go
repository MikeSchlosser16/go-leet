func searchInsert(nums []int, target int) int {
    for i, num := range nums{
        // if it matches, return
        // if its greater, return i as we'd insert here before incrementing
        if num == target || num > target{
            return i
        }
        // Handle insert at the end case
        if i == len(nums) - 1{
            return len(nums)
        }
    }
    return -1
}


// Obvious binary search problem, array sorted..

func searchInsert(nums []int, target int) int {
    return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, left int, right int) int{
    if left >= right {
        if target > nums[left] {
            return left+1
        } else {
            return left
        }
    } else {
        mid := int((left+right)/2)
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            right = mid -1
        } else {
            left = mid + 1
        }
        return binarySearch(nums, target, left, right)
    }
}
