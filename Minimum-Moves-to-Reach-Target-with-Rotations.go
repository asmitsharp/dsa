1func minimumMoves(grid [][]int) int {
2   n := len(grid)
3
4   type State struct {
5    r, c, dir int
6   }
7
8   queue := []State{{0,0,0}}
9   visited := make([][][]bool, n)
10
11   for i := range visited {
12    visited[i] = make([][]bool, n)
13    for j := range visited[i] {
14        visited[i][j] = make([]bool, 2)
15    }
16   }
17
18   visited[0][0][0] = true
19   steps := 0
20
21   for len(queue) > 0 {
22    size := len(queue)
23
24    for i := 0; i < size; i++ {
25        curr := queue[0]
26        queue = queue[1:]
27
28        r,c, dir := curr.r, curr.c, curr.dir
29
30        if r == n-1 && c == n-2 && dir == 0 {
31            return steps
32        }
33
34        if dir == 0 {
35            if c+2 < n && grid[r][c+2] == 0 && !visited[r][c+1][0] {
36                visited[r][c+1][0] = true
37                queue = append(queue, State{r, c+1, 0})
38            }
39
40            if r+1 < n && 
41                grid[r+1][c] == 0 &&
42                grid[r+1][c+1] == 0 && 
43                !visited[r+1][c][0] {
44                    visited[r+1][c][0] = true
45                    queue = append(queue, State{r + 1, c, 0})
46                }
47            
48            if r+1 < n && 
49                grid[r+1][c] == 0 && 
50                grid[r+1][c+1] == 0 &&
51                !visited[r][c][1] {
52                    visited[r][c][1] = true
53                    queue = append(queue, State{r,c,1})
54                }
55        } else {
56             if r+2 < n && grid[r+2][c] == 0 && !visited[r+1][c][1] {
57                    visited[r+1][c][1] = true
58                    queue = append(queue, State{r + 1, c, 1})
59                }
60
61                // move right
62                if c+1 < n &&
63                    grid[r][c+1] == 0 &&
64                    grid[r+1][c+1] == 0 &&
65                    !visited[r][c+1][1] {
66
67                    visited[r][c+1][1] = true
68                    queue = append(queue, State{r, c + 1, 1})
69                }
70
71                // rotate anti-clockwise
72                if c+1 < n &&
73                    grid[r][c+1] == 0 &&
74                    grid[r+1][c+1] == 0 &&
75                    !visited[r][c][0] {
76
77                    visited[r][c][0] = true
78                    queue = append(queue, State{r, c, 0})
79                }
80        }
81    }
82    steps++
83   }
84   return -1
85}
86