1func lengthOfLongestSubstring(s string) int {
2    char := make(map[byte]int)
3    left := 0
4    maxLen := 0
5
6    for right := 0; right < len(s); right++ {
7        char[s[right]]++
8
9        for char[s[right]] > 1 {
10            char[s[left]]--
11            left++
12        }
13
14        windowLen := right - left + 1
15
16        if windowLen > maxLen {
17            maxLen = windowLen
18        }
19    }
20
21    return maxLen
22}