// bad
class Solution {
    public int firstUniqChar(String s) {

        for (int i = 0; i < s.length(); i++){
            boolean unique = true;
            for (int j = i + 1; j < s.length(); j++){
                if (s.charAt(i) == s.charAt(j)){
                    unique = false;
                    break;
                }
            }
            if(unique){
                return i;
            }
        }
        return -1;

    }
}

// better
class Solution {
    public int firstUniqChar(String s) {
        int[] counts = new int[26]; // array rather than dict to keep ordering

        // store at the character index a count.. ie counts[0] = # 'a', counts[1] = # 'b'
        for (char ch : s.toCharArray()){
            counts[ch - 'a'] = counts[ch - 'a'] + 1;
        }

        // Iterate over s, ie. "leetcode"
        // counts[l] == 1, return 0
        for (int i = 0; i < s.length(); i++){
            if (counts[s.charAt(i) - 'a'] == 1){
                return i;
            }
        }
        return -1;
    }
}
