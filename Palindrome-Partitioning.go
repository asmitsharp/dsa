1func partition(s string) [][]string {
2    result := [][]string{}
3    path := []string{}
4    var backtrack func(start int)
5    backtrack = func(start int) {
6        if start == len(s) {
7            temp := make([]string, len(path))
8            copy(temp, path)
9            result = append(result, temp)
10            return
11        } 
12
13        for i := start; i < len(s); i++ {
14            if isPalindrome(s[start:i+1]) {
15                path = append(path, s[start:i+1])
16                backtrack(i+1)
17                path = path[:len(path)-1]
18            }
19        }
20    }
21    backtrack(0)
22    return result
23}
24
25func isPalindrome(s string) bool {
26    l, r := 0, len(s)-1
27    for l < r {
28        if s[l] != s[r] {
29            return false
30        }
31        l++
32        r--
33    }
34    return true
35}