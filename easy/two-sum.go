// Brute force - iterate though the list and then using the rest of the list see if the value matches, return the indexes
// Time complexity: O(n^2) - for each element in the list n, loop over the rest of the array n = n * n = n^2
// Space complexity = O(1)
func twoSum(nums []int, target int) []int {
    for i, num := range nums {
        shortenedList := nums[i+1:]
        for j, num2 := range shortenedList{
            if num + num2 == target {
                return []int{i, j+i+1}
            }
        }
    }
    return nil
}

// Put in a map, loop over the map and look for the complementing int to make target
// Time complexity: O(n) - iterater over list twice, but search in map is O(1) so better
// Space complexity = O(n) - need to store all n numbers in the map
func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int) // K: num value V: index
    for i, num := range nums {
        numsMap[num] = i
    }
    for num, index := range numsMap{
        seek := target - num
        seekIndex := numsMap[seek]
        if seekIndex != 0 && seekIndex != index {
            return []int{seekIndex, index}
        }
    }
    return nil
}

// One pass though the list using a map..
// Instead of appending to the map and THEN iterating, we can check if its in the map and then add if not on the same pass, and essentially look back
// Time complexity: O(n) - loop over the list of n length once
// Space complexity: O(n) - map could store up to n elements
func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int)
    for index, num := range nums {
        seek := target - num
        if _, found := numsMap[seek]; found { // cant check if == 0 here beacuse the index may be zero!
            return []int{numsMap[seek], index}
        }
        numsMap[num] = index
    }
    return nil
}
