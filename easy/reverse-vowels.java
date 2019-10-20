// Obvious two pointer problem
// One at the start, one at the end
class Solution {
    public String reverseVowels(String s) {
        String vowels = "aeiouAEIOU";
        char[] chars = s.toCharArray();
        int start = 0;
        int end = s.length() - 1;

        // Go until half way
        while (start < end){

            // Keep searching for the next vowel at the start. Check for out of bounds
            while ((start < end) && !vowels.contains(chars[start]+"")){
                start++;
            }

            // Once we found a vowel at the start, find the one next with the second pointer from the end
            while ((start < end) && !vowels.contains(chars[end]+"")){
                end--;
            }

            // Once here, start and end are both on index of vowels, swap them
            char temp = chars[start];
            chars[start] = chars[end];
            chars[end] = temp;

            start++;
            end--;

        }

        return new String(chars); // Convert []char back to string to avoid immutable issues
    }
}