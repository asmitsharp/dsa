func groupAnagrams(strs []string) [][]string {
    agroup := make(map[string][]string)

    for _, str := range strs {
        key := key(str)
        agroup[key] = append(agroup[key], str)
    }

     result := make([][]string, 0, len(agroup))
     for _, group := range agroup {
        result = append(result, group)
     }

     return result

}

func key(str string) string {
    charCount := make([]int, 26)

    for _, char := range str {
        charCount[char - 'a']++
    }

    var keyBuilder strings.Builder
    for i, count := range charCount {
        if count > 0 {
            keyBuilder.WriteByte(byte('a' + i))
            keyBuilder.WriteString(strconv.Itoa(count))
        }
    }

    return keyBuilder.String()
}