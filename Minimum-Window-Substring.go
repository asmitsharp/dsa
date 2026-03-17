1func minWindow(s string, t string) string {
2    left := 0
3    need := make(map[byte]int)
4    win := make(map[byte]int)
5    minLen := len(s) + 1
6    minStart := 0
7
8    for i := 0; i < len(t); i++ {
9        need[t[i]]++
10    }
11
12    var match int
13    required := len(need)
14
15    for right := 0; right < len(s); right++ {
16        c := s[right]
17        win[c]++
18
19        if win[c] == need[c] {
20            match++
21        }
22
23        for match == required {
24            winLen :=  right - left + 1
25            if winLen < minLen {
26                minLen = winLen
27                minStart = left
28            }
29
30            leftChar := s[left]
31            win[leftChar]--
32            if win[leftChar] < need[leftChar]{
33                match--
34            }
35            left++
36        }
37    }
38     if minLen == len(s) + 1 {
39        return ""
40    }
41    return s[minStart : minStart+minLen]
42}