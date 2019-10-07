// Split, check len(n), len(n-1)... to find first valid
import "unicode"
import "fmt"

func lengthOfLastWord(s string) int {

    blocks := strings.Split(s, " ")

    if len(blocks) == 1{
         if isValidWord(blocks[0]){
            return len(blocks[0])
        }
    }

    for i := len(blocks)-1; i >= 0; i--{
        if isValidWord(blocks[i]){
            return len(blocks[i])
        }
    }
    return 0
}

func isValidWord(s string) bool {
    if s == ""{
        return false
    }

    word := true
    hasChar := false
    for _, ch := range s{
        if !unicode.IsLetter(rune(ch)){
            word = false
            break
        }else{
            hasChar = true

        }
    }
    return word && hasChar
}






// Simple to get most, but fails on "a " for example
func lengthOfLastWord(s string) int {
    words := strings.Split(s, " ")
    if len(words) < 1{
        return 0
    }
    return len(words[len(words)-1])
}
