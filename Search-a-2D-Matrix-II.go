1func searchMatrix(matrix [][]int, target int) bool {
2    m := len(matrix)
3    n := len(matrix[0])
4
5    row := 0
6    col := n-1
7
8    for row < m && col >= 0 {
9        if matrix[row][col] == target {
10            return true
11        }
12        
13        if target < matrix[row][col] {
14            col--
15        } else {
16            row++
17        }
18    }
19
20    return false
21}