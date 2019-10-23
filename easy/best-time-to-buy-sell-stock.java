class Solution {
    public int maxProfit(int[] prices) {
        int currentMax = 0;
        for(int i = 0; i < prices.length; i++){
            for(int j = i + 1; j < prices.length; j++){
                System.out.println(i-j);
                if((prices[j] - prices[i]) > currentMax){
                    currentMax = prices[j] - prices[i];
                }
            }
        }
        return currentMax;
        }
}
