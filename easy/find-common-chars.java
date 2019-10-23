// Keep a "common" map that contains the least match
// Create a temp map for each loop though, then compare to common map. If less, must use that value
class Solution {
    public List<String> commonChars(String[] A) {
        List<String> result = new ArrayList<>();
        int[] charCount = new int[26]; // problem states all chars are lowercase letters. This is the "common" dict

        for(int i = 0; i < A[0].length(); i++){
            charCount[A[0].charAt(i) - 'a']++;
        }

        for (int j = 1; j < A.length; j++){ // The rest of the words
            int[] currentWordCharCount = new int[26];
            for (int k = 0; k < A[j].length(); k++){ // The characters in each word
                currentWordCharCount[A[j].charAt(k) - 'a']++;
            }

            // Update the "common" dict
            for (int l = 0; l < 26; l++){
                if (currentWordCharCount[l] < charCount[l]){ // If its less, this is now the "bottom" bound
                    charCount[l] = currentWordCharCount[l];
                }
            }
        }

        // Append to result to return
        for( int i = 0; i < 26; i++){ // Check each letter
            for(int j = 0; j < charCount[i]; j++){ // We need to add duplicates, potentially..
                result.add(Character.toString((char) ('a' + i)));
            }
        }
        return result;
    }
}
