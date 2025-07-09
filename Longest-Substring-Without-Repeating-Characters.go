func lengthOfLongestSubstring(s string) int {
    maxL := 0

    wStart := 0
    charIndex := make(map[byte]int)

    for i := 0 ; i < len(s); i++ {

        if lastIndex, found := charIndex[s[i]]; found && lastIndex >= wStart {
            wStart = lastIndex + 1
        }

        charIndex[s[i]] = i
        currL := i - wStart + 1
        if currL > maxL {
            maxL = currL
        }
    }

    return maxL
}