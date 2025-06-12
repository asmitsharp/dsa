func isAnagram(s string, t string) bool {
    // if len(s) != len(t) {
    //     return false
    // }

    // charCount := make(map[rune]int)

    // for _, char := range s {
    //     charCount[char]++
    // }

    // for _, char := range t {
    //     charCount[char]--
    //     if charCount[char] < 0 {
    //         return false
    //     }
    // }

    // for _ , count := range charCount {
    //     if count != 0 {
    //         return false
    //     }
    // }

    // return true

    chars := make([]int , 26)

    for _, v := range s {
        chars[v - 'a']++
    }

    for _, v := range t {
        chars[v - 'a']--
    }

    for _, v := range chars {
       if v != 0 {
\t\t\treturn false
\t\t} 
    }

    return true
}