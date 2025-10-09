func reverseOnlyLetters(s string) string {
    chars := []rune(s)
    l, r := 0, len(s) - 1

    for  l < r {
        for l < r && !unicode.IsLetter(chars[l]) {
            l++
        }

        for l < r && !unicode.IsLetter(chars[r]) {
            r--
        }

        if l < r {
            chars[l], chars[r] = chars[r], chars[l]
        }

        l++
        r--
    }

    return string(chars)
}