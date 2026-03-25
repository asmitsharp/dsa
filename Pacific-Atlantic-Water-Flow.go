1func pacificAtlantic(heights [][]int) [][]int {
2    m := len(heights)
3    n := len(heights[0])
4
5    pac := make([][]bool, m)
6    atl := make([][]bool, m)
7
8    for i := range pac {
9        pac[i] = make([]bool, n)
10        atl[i] = make([]bool, n)
11    }
12
13    var dfs func(r, c int, visit [][]bool) 
14    dfs = func(r, c int, visit [][]bool) {
15        visit[r][c] = true 
16
17        dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
18
19        for _, d := range dirs {
20            nr := r + d[0]
21            nc := c + d[1]
22
23            if nr < 0 || nc < 0 || nr >= m || nc >= n {
24                continue
25            }
26
27            if visit[nr][nc] {
28                continue
29            }
30
31            if heights[nr][nc] < heights[r][c] {
32                continue
33            }
34
35            dfs(nr, nc, visit)
36        }
37    }
38
39    res := [][]int{}
40
41    for i := 0; i < m; i++ {
42        dfs(i, 0, pac)
43        dfs(i, n-1, atl)
44    }
45
46    for i := 0; i < n; i++ {
47        dfs(0, i, pac)
48        dfs(m-1, i, atl)
49    }
50
51    for i := 0;  i < m; i++ {
52        for j := 0;  j < n;  j++ {
53            if pac[i][j] && atl[i][j] {
54                res = append(res, []int{i,j})
55            }
56        }
57    }
58
59    return res
60}