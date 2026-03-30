1func generateParenthesis(n int) []string {
2    res := []string{}
3    backtrack("", 0, 0, n, &res)
4    return res
5}
6
7func backtrack(path string, open, close, n int, res *[]string) {
8    if len(path) == 2*n {
9        *res = append(*res, path)
10        return
11    }
12
13    if open < n {
14        backtrack(path+"(", open+1, close, n, res)
15    }
16
17    if close < open {
18        backtrack(path+")", open, close+1, n, res)
19    }
20
21}