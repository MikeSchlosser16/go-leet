import java.util.*;

// did in java for stack impl, trivial to do in go with a slice
class Solution {
    public boolean isValid(String s) {
        if (s == ""){
            return false;
        }

        HashMap<Character, Character> inverses = new HashMap();
        inverses.put('}', '{');
        inverses.put(')', '(');
        inverses.put(']', '[');

        Stack<Character> stack = new Stack();

        // Classic stack problem
        // if its an opening, push to stack.. the next non opening should close it or its invalid
        for (char ch : s.toCharArray()){
            if (isOpening(ch)){
                stack.push(ch);
            }else{
                 if (stack.size() != 0){
                    Character stackCh = stack.peek();
                    if(inverses.get(ch) == stackCh){
                        stack.pop(); // pop it, valid closing
                    }else{
                        return false;
                    }
                }else{
                     return false;
                 }
            }
        }

       return stack.size() == 0;
    }

    public boolean isOpening(char ch) {
        return ch == '{' || ch == '(' || ch == '[';
    }

}
