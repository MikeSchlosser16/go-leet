// Dynamic programming problem - need to notice Fibonacci part of this 
func climbStairs(n int) int {
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    for i:= 2; i <= n; i++{
        // We know we can only take 2 steps at most
        // We MUST have came from the step before or two steps before it
        // Getting to the ith step is basially where could you have came from?
        // We can only get there from prior step or 2 prior, and we have already solved those sub problems
        dp[i] = dp[i-1] + dp[i-2]
    }

    return dp[n]
}
