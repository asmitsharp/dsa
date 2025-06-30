func firstUniqChar(s string) int {
    // freq := make(map[rune]int)
    // for _, char := range s {
    //     freq[char]++
    // }

    // queue := make([]int, 0)
    // for i := range s {
    //     queue = append(queue, i)
    // }

    // for len(queue) > 0 {
    //     index := queue[0]

    //     char := rune(s[index])

    //     for freq[char] == 1 {
    //         return index
    //     }

    //     queue = queue[1:]
    // }
    // return -1

    count := make([]int, 26)
    for _, char := range s {
        count[char-'a']++
    }

    for i, ch := range s {
        if count[ch-'a'] == 1 {
            return i
        }
    }

    return -1
}