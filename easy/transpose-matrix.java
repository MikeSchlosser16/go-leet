class Solution {
    public int[][] transpose(int[][] A) {

        int columns = A.length;
        int rows = A[0].length;

        int[][] result = new int[rows][columns];

        for(int i = 0; i < rows; i++){
            for (int j = 0; j < columns; j++){
                result[i][j] = A[j][i];
            }
        }

        return result;

    }
}