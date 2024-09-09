func makeEqual(words []string) bool {
    n := len(words)
    m := make(map[rune]int)

    for _, word := range words {
        for _, char := range word {
            m[char]++
        }
    }

    for _, count := range m {
        if count % n != 0 {
            return false
        }
    }

    return true
}