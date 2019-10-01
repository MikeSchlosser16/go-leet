/*
Convert roman string to int value
Important thing to note is that if the a characters right part is bigger than itself, we need to subtract.
For example, IV = 5 - 1 since V > I ---> 5 > 1 ---> 5-1, since right most is greater
Otherwise, we can add, VI ---> I < V -----> 5 + 1 since right most is smaller
*/

func romanToInt(s string) int {

    valueBySymbol := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    sum := valueBySymbol[rune(s[len(s) - 1])] // Initial value is the right char of string

    // Read right to left, starting at second char
    // If right > left, add
    // If right < left, subtract
    for i := len(s) - 2; i >= 0; i--{
      if valueBySymbol[rune(s[i])] >= valueBySymbol[rune(s[i+1])]{
        sum += valueBySymbol[rune(s[i])]
      } else {
        sum -= valueBySymbol[rune(s[i])]
      }
    }
    return sum
}
