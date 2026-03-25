1func orangesRotting(grid [][]int) int {
2    m := len(grid)
3    n := len(grid[0])
4
5    queue := [][]int{}
6    fresh := 0
7
8    for i := 0; i < m; i++ {
9        for j := 0; j < n; j++ {
10            if grid[i][j] == 2 {
11                queue = append(queue, []int{i, j})
12            }
13            if grid[i][j] == 1 {
14                fresh++
15            }
16        }
17    }
18
19    minutes := 0
20    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
21
22    for len(queue) > 0 && fresh > 0 {
23         size := len(queue)
24
25         for i := 0; i < size; i++ {
26            cell := queue[0]
27            queue = queue[1:]
28
29            r, c := cell[0], cell[1]
30
31            for _, d :=  range dirs {
32                nr := r + d[0]
33                nc := c + d[1]
34
35
36                if nr < 0 || nc < 0 || nr >= m || nc >= n {
37                    continue
38                } 
39
40                if grid[nr][nc] != 1 {
41                    continue
42                }
43
44                grid[nr][nc] = 2
45                fresh--
46                queue = append(queue, []int{nr, nc})
47            }
48         }
49         minutes++
50    }
51
52    if fresh > 0 {
53        return -1
54    }
55
56    return minutes
57}