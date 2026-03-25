1func exist(board [][]byte, word string) bool {
2    m := len(board)
3    n := len(board[0])
4
5    var dfs func(r, c, index int) bool
6
7    dfs = func(r, c, index int) bool {
8        if len(word) == index {
9            return true
10        }
11
12        if r < 0 || c < 0 || r >= m || c >= n {
13            return false
14        }
15
16        if board[r][c] != word[index] {
17            return false
18        }
19
20        temp := board[r][c]
21        board[r][c] = '#'
22
23        found := dfs(r+1, c , index+1) || 
24                    dfs(r-1, c , index+1) ||
25                    dfs(r, c+1 , index+1) ||
26                    dfs(r, c-1 , index+1)
27        
28        board[r][c] = temp
29
30        return found
31
32    }
33
34    for i := 0; i < m; i++ {
35        for j := 0; j < n; j++ {
36            if dfs(i, j, 0) {
37                return true
38            }
39        }
40    }
41
42    return false
43
44}