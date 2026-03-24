1func countNegatives(grid [][]int) int {
2    count := 0
3    row := len(grid) - 1
4    col := 0
5
6    for row >= 0 && col < len(grid[0]) {
7        if grid[row][col] < 0 {
8            count += len(grid[0]) - col
9            row--
10        } else {
11            col++
12        }
13    } 
14    return count
15}