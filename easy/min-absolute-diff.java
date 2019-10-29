// Much better, using a pointer and checking its value and the element prior
class Solution {
    public List<List<Integer>> minimumAbsDifference(int[] arr) {
        Arrays.sort(arr); // Sort in place
        List<List<Integer>> ans = new ArrayList<>();
        int diff = Integer.MAX_VALUE; // Initialize starting diff value

        //start looping over array to find real min element. Each time we found smaller difference
        //we reset resulting list and start building it from scratch. If we found pair with the same
        //difference as min - add it to the resulting list
        for(int i = 1; i < arr.length; i++){
            if(arr[i] - arr[i-1] < diff){ // sorted so this is fine without worrying about abs
                diff = arr[i] - arr[i-1];
                ans.clear();
            }
            ans.add(Arrays.asList(arr[i-1], arr[i]));
        }
        return ans;
    }
}


// Really bad
class Solution {
    public List<List<Integer>> minimumAbsDifference(int[] arr) {
        List<List<Integer>> result = new ArrayList<>();

        // Find mininium abs difference
        int minDiff = Math.abs(arr[0] - arr[1]);
         for (int i = 0; i < arr.length; i++){
            for(int j = i + 1; j < arr.length; j++){
                int currentDiff = Math.abs(arr[i] - arr[j]);
                if (currentDiff < minDiff){
                    minDiff = currentDiff;
                }
            }
        }

        for (int i = 0; i < arr.length; i++){
            for(int j = i + 1; j < arr.length; j++){
                if (Math.abs(arr[i] - arr[j]) == minDiff){
                    List<Integer> sorted = new ArrayList<>();
                    sorted.add(Math.min(arr[i], arr[j]));
                    sorted.add(Math.max(arr[i], arr[j]));
                    int elementSum = arr[i] + arr[j];
                    if(result.size() == 0){
                        result.add(0, sorted);
                    }else{

                        for (int k = 1; k < result.size() + 1; k++){
                                int currentSum = result.get(k-1).get(0) + result.get(k-1).get(1);
                                if (elementSum < currentSum){
                                    result.add(k-1, sorted);
                                    break;
                                }
                        }
                    }
                }
            }
        }
        return result;
    }
}
