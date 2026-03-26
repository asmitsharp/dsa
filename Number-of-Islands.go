1// func numIslands(grid [][]byte) int {
2//     row := len(grid)
3//     cols := len(grid[0])
4//     count := 0
5
6//     for r := 0; r < row; r++ {
7//         for c := 0; c < cols; c++ {
8//             if grid[r][c] == '1' {
9//                 count++
10//                 dfs(grid, r, c)
11//             }
12//         }
13//     }
14
15//     return count
16// }
17
18// func dfs(grid [][]byte, r, c int) {
19//     if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
20//         return
21//     }
22
23//     grid[r][c] = '0'
24//     dfs(grid, r+1, c)
25//     dfs(grid, r-1, c)
26//     dfs(grid, r, c+1)
27//     dfs(grid, r, c-1)
28// }
29
30
31func numIslands(grid [][]byte) int {
32    rows := len(grid)
33    cols := len(grid[0])
34
35    dirs := [][]int{
36        {1,0},
37        {-1,0},
38        {0,1},
39        {0,-1},
40    }
41
42    count := 0
43
44    for r := 0; r < rows; r++ {
45        for c := 0; c < cols; c++ {
46
47            if grid[r][c] == '1' {
48                count++
49
50                queue := [][]int{{r,c}}
51                grid[r][c] = '0'
52
53                for len(queue) > 0 {
54                    cell := queue[0]
55                    queue = queue[1:]
56
57                    cr, cc := cell[0], cell[1]
58
59                    for _, d := range dirs {
60                        nr := cr + d[0]
61                        nc := cc + d[1]
62
63                        if nr >= 0 && nr < rows &&
64                           nc >= 0 && nc < cols &&
65                           grid[nr][nc] == '1' {
66
67                            grid[nr][nc] = '0'
68                            queue = append(queue, []int{nr,nc})
69                        }
70                    }
71                }
72            }
73        }
74    }
75
76    return count
77}
78