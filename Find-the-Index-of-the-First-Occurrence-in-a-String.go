func lps(needle string) []int {
    n := len(needle)
    lps := make([]int, n)

    length := 0
    i := 1

    for i < n {
        if needle[i] == needle[length] {
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

func strStr(haystack string, needle string) int {
   if len(needle) == 0 {
        return 0  
    }
    if len(haystack) < len(needle) {
        return -1
    }

    lps := lps(needle)
    i := 0 // text
    j := 0 // pattern


    for i < len(haystack) {
        if haystack[i] == needle[j] {
            i++
            j++
        }

        if j == len(needle) {
            return i - j 
        } else if i < len(haystack) && haystack[i] != needle[j] {
            if j != 0{
                j = lps[j-1]
            } else {
                i++
            }
        }
    }

return -1
}