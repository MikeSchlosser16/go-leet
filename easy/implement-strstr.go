func strStr(haystack string, needle string) int {
    if needle == ""{
        return 0
    }

    lenHaystack := len(haystack)
    lenNeedle := len(needle)

    if lenHaystack < lenNeedle{
        return -1
    }

    // Optimization - we can't go past lenHaystack-lenNeedle + 1, since if we do there isnt room for the needle left to possibly match
    for i := 0; i < lenHaystack-lenNeedle + 1; i++{
        found := true
        // iterate through needle chars, if one doesnt match break and go to the next start in the haystack
        for j := 0; j < lenNeedle; j++{
            // notice i+j - need to check at the starting element of haystack that matched(i) + the next, next +1 etc (j)
            if haystack[i+j] != needle[j]{
                found = false
                break
            }
        }
        if found{
            return i
        }
    }
  return -1
}
