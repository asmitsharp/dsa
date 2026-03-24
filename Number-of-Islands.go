1func numIslands(grid [][]byte) int {
2    row := len(grid)
3    cols := len(grid[0])
4    count := 0
5
6    for r := 0; r < row; r++ {
7        for c := 0; c < cols; c++ {
8            if grid[r][c] == '1' {
9                count++
10                dfs(grid, r, c)
11            }
12        }
13    }
14
15    return count
16}
17
18func dfs(grid [][]byte, r, c int) {
19    if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != '1' {
20        return
21    }
22
23    grid[r][c] = '0'
24    dfs(grid, r+1, c)
25    dfs(grid, r-1, c)
26    dfs(grid, r, c+1)
27    dfs(grid, r, c-1)
28}