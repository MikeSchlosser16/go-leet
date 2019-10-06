import "math"

func maxSubArray(nums []int) int {
	max := math.MinInt32 // handle default val being zero
	for i := 0; i < len(nums); i++ {
		internalMax := math.MinInt32
		currentVal := 0
		for j := i; j < len(nums); j++ { // get all possible values at index 0, 1, 2..
			currentVal += nums[j]
			if currentVal > internalMax { // get 'internal max' for each element
				internalMax = currentVal
			}
		}
		if internalMax > max { // if this elements max > current max, replace
			max = internalMax
		}
	}
	return max
}



// Dynamic programming
// Take the best of the previous element, check if we just take the value or extend to what we did before
// [-2, 1, -3, 4] input
// [-2,  ,   ,  ] -2 best we can do at 1
// [-2, 1,   ,  ] do we take 1, or extent the prior? max(1,1-2) == 1
// [-2, 1, -2,  ] do we take -3, or extent the prior? max(-3,-3+1) == 1
// [-2, 1, -2 , 4] do we take 4, or extent the prior? max(4,4-2) == 1
// Then we take the max of the possibilities, which is 4
class LinearTimeSolution {

  public int maxSubArray(int[] nums) {
    int maxSoFar = nums[0];
    int maxEndingHere = nums[0];

    for (int i = 1; i < nums.length; i++){

      maxEndingHere = Math.max(maxEndingHere + nums[i], nums[i]);
      maxSoFar = Math.max(maxSoFar, maxEndingHere);
    }

    return maxSoFar;
  }

}
