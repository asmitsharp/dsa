func maxLengthBetweenEqualCharacters(s string) int {
    ans := -1
    last := make(map[rune]int)

    for j, c := range s {
        if i, ok := last[c]; ok {
            tmp := j - i - 1
            if tmp > ans {
                ans = tmp
            }
        } else {
            last[c] = j
        }
    }
    return ans
}