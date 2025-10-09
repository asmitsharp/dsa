import "unicode"

func isPalindrome(s string) bool {
    var str []rune

    for _, i := range s {
        if unicode.IsLetter(i) || unicode.IsDigit(i) {
            str = append(str, unicode.ToLower(i))
        }
    }

    l, r := 0, len(str) - 1
    for l < r {
        if str[l] != str[r] {
            return false
        }
        l++
        r--
    }

    return true
}
