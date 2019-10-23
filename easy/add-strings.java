class Solution {
    public String addStrings(String num1, String num2) {

        StringBuilder sb = new StringBuilder();

        int num1Counter = num1.length() - 1;
        int num2Counter = num2.length() - 1;

        int carry = 0;

        while (num1Counter >= 0 || num2Counter >= 0){
            int n1 = 0;
            int n2 = 0;
            if(num1Counter >= 0){
                 n1 = num1.charAt(num1Counter) - '0';
            }
            if(num2Counter >= 0){
                n2 = num2.charAt(num2Counter) - '0';
            }


            int current = (n1 + n2 + carry) % 10;
            sb.insert(0, current);
            carry = (n1 + n2 + carry) / 10;

            num1Counter--;
            num2Counter--;
        }

        if(carry != 0){
            sb.insert(0, 1);
        }

        return sb.toString();
    }
}