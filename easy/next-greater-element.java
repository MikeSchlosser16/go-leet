class Solution {
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        HashMap<Integer, Integer> next_greatest = new HashMap();
        Stack<Integer> stack = new Stack();

        // Loop thru array to build map with next greatest
        for (Integer num : nums2){
            // when stack is empty, put num on stack
            while (!stack.isEmpty() && stack.peek() < num){
                next_greatest.put(stack.pop(), num);
            }
            stack.push(num);
        }
        for(int i = 0; i < nums1.length; i++){
            nums1[i] = next_greatest.getOrDefault(nums1[i], -1);
        }
        return nums1;
    }
}



/*
stack = [1]
is 1 < 3, yes, we found next greatest, 3. pop 1 off, put (1,3), so 1's ext greatest is 3.
stack = [3]
3 < 4, map.put(3,4) pop 3
stack = [4]
*/
