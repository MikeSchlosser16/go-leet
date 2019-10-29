public static List<String> findWords(String[] words) {

        List<String> result = new ArrayList<String>();

        String row1 = "qwertyuiop";
        String row2 = "asdfghjkl";
        String row3 = "zxcvbnm";

        for (int i = 0; i < words.length; i++){
            String word = words[i];
            boolean valid = true;
            String initial = Character.toString(word.charAt(0)).toLowerCase();

            System.out.println("initial: " + initial);

            String toUse = "";

            if(row1.contains(initial)){
                toUse = row1;
            }else if(row2.contains(initial)){
                toUse = row2;
            }else{
                toUse = row3;
            }

            System.out.println("to use: " + toUse);
            for(int j = 1; j < word.length(); j++){
                if (!toUse.contains(Character.toString(word.charAt(j)).toLowerCase())){
                    valid = false;
                    break;
                }
            }
            if(valid){
                result.add(word);
            }
        }
        return result;
    }