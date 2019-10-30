// BFS - Wide, use a queue (FIFO), go out level by level
// DFS - Deep, using a stack (recursion, LIFO)


/*
  For string "abc"

           a
        /             \
      a                 A
     /\                /\
  ab   aB           Ab       AB
  /\     /\        /\         /\
abc abC aBc aBC   Abc AbC   ABc ABC


Imagine a tree, we start from the dummy node, left node means choice 1(lower) and the right means choice 2(upper).




For BFS "ab"
              queue:      i:    size:       j:    cur:      offers:
              "ab"        0     1           0     "ab"      Ab, ab
              "Ab, ab"    1     2           0     "Ab"      [Ab, AB
                                            1     "ab"      ab, aB]
*/

// BFS
class Solution {
    public List<String> letterCasePermutation(String S) {
        if (S == null){
            return new LinkedList<>();
        }
        // Usually we would keep track of "seen"
        Queue<String> queue = new LinkedList<>(); // Use queue for FIFO
        queue.offer(S); // Add the "start" node, or in this case string

        for(int i = 0; i < S.length(); i++){
            if(Character.isDigit(S.charAt(i))){
                continue;
            }
            int size = queue.size();
            for(int j = 0; j < size; j++){ // While queue is not empty
                String cur = queue.poll(); // Pull a node
                char[] chs = cur.toCharArray();
                chs[i] = Character.toUpperCase(chs[i]); // Conver this char to upper
                queue.offer(String.valueOf(chs)); // Put back in queue, adding "unseen"
                chs[i] = Character.toLowerCase(chs[i]); // Convert this char to lower
                queue.offer(String.valueOf(chs));  // Put back in queue , adding "unseen"
            }
        }
        return new LinkedList<>(queue);
    }
}

// DFS
class Solution {
    public List<String> letterCasePermutation(String S) {
        if (S == null) {
            return new LinkedList<>();
        }

        List<String> res = new LinkedList<>();
        helper(S.toCharArray(), res, 0);
        return res;
    }

    public void helper(char[] chs, List<String> res, int pos) {
        if (pos == chs.length) {
            res.add(new String(chs));
            return;
        }
        if (chs[pos] >= '0' && chs[pos] <= '9') {
            helper(chs, res, pos + 1);
            return;
        }

        chs[pos] = Character.toLowerCase(chs[pos]);
        helper(chs, res, pos + 1);

        chs[pos] = Character.toUpperCase(chs[pos]);
        helper(chs, res, pos + 1);
    }
}a
