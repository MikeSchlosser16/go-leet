
public int maxNumberOfBalloons(String text) {
  int[] charCounts = new int[26]; // 26 lowercase letters
  for(int i = 0; i < text.length(); i++){
    charCounts[text.charAt(i) - 'a']++; // conver to char, subtracting ascii value
  }
  int min = charCounts[1]; // b
  min = Math.min(min, charCounts[0]); // a
  min = Math.min(min, charCounts[11] / 2); // l
  min = Math.min(min, charCounts[14] / 2); // o
  min = Math.min(min, charCounts[13]); // n
  return min;
}


// Bad
class Solution {
    public int maxNumberOfBalloons(String text) {
        HashMap<Character, Integer> charCount = new HashMap<>();
        for (Character ch : text.toCharArray()){
             if(charCount.containsKey(ch)){
                 int count = charCount.get(ch);
                 charCount.put(ch, count + 1);
            }else{
                charCount.put(ch, 1);
            }
        }

       // do we have all chars?
        if (!charCount.containsKey('b') ||  !charCount.containsKey('a') || !charCount.containsKey('l') ||                   !charCount.containsKey('o') || !charCount.containsKey('n')){
            return 0;
        }

        // divide l and o by 2
        charCount.put('l', charCount.get('l') / 2);
        charCount.put('o', charCount.get('o') / 2);


        int smallest = Integer.MAX_VALUE;
        if (charCount.get('b') < smallest ) smallest = charCount.get('b') ;
        if (charCount.get('a') < smallest ) smallest = charCount.get('a') ;
        if (charCount.get('l') < smallest ) smallest = charCount.get('l') ;
        if (charCount.get('o') < smallest ) smallest = charCount.get('o') ;
        if (charCount.get('n') < smallest ) smallest = charCount.get('n') ;


        return smallest;
    }
}
