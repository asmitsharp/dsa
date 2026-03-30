1func letterCasePermutation(s string) []string {
2    res := []string{}
3    path := []byte(s)
4
5    var backtrack func(index int)
6    backtrack = func(index int) {
7        if len(path) == index {
8            res = append(res, string(path))
9            return
10        }
11
12        backtrack(index+1)
13
14        if path[index] >= 'a' && path[index] <= 'z' {
15            path[index] = path[index] - 'a' + 'A'
16            backtrack(index+1)
17            path[index] = path[index] - 'A' + 'a'
18        } else if path[index] >= 'A' && path[index] <= 'Z' {
19            path[index] = path[index] - 'A' + 'a'
20            backtrack(index+1)
21            path[index] = path[index] - 'a' + 'A'
22        }
23    }
24
25    backtrack(0)
26    return res
27}
28