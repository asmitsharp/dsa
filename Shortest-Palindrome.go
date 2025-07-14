func shortestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }


    reverse := reverseString(s)
  combined := s + "#" + reverse
    
    // Build LPS array for the combined string
    lps := buildLPS(combined)

    overlap := lps[len(combined) - 1]

       return reverse[:len(s)-overlap] + s


}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func buildLPS(pattern string) []int {
    n := len(pattern)
    lps := make([]int, n)
    
    length := 0
    i := 1
    
    for i < n {
        if pattern[i] == pattern[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]
            } else {
                lps[i] = 0
                i++
            }
        }
    }
    
    return lps
}