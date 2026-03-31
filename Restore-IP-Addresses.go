1func restoreIpAddresses(s string) []string {
2    result := []string{}
3    
4    var backtrack func(path string, start int, parts int)
5    backtrack = func(path string, start int, parts int) {
6        if parts == 4 && start == len(s) {
7            result = append(result, path[:len(path)-1])
8            return
9        }
10
11        if parts == 4 || start == len(s) {
12            return
13        }
14
15        for i := start; i < len(s) && i < start+3; i++ {
16            segment := s[start:i+1]
17
18            if !isValid(segment) {
19                continue
20            }
21
22            backtrack(path+segment+".", i+1, parts+1)
23        }
24    }
25    backtrack("",0,0)
26    return result
27}
28
29func isValid(s string) bool {
30    if len(s) > 1 && s[0] == '0' {
31        return false
32    }
33
34    num := 0
35    for i := 0; i < len(s); i++ {
36        num = num*10 + int(s[i]-'0')
37    }
38
39    return num <= 255
40}