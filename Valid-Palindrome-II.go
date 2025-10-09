func validPalindrome(s string) bool {
    l,r := 0, len(s) - 1
    for l < r {
        if s[l] != s[r] {
            check1 := valid(s, l+1, r)
            check2 := valid(s, l, r-1)

            return check1 || check2
        }

        l++
        r--
    }
    return true
}

func valid(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}