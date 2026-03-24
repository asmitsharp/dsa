1func transpose(matrix [][]int) [][]int {
2    m := len(matrix)      // original rows
3    n := len(matrix[0])   // original cols
4
5    // result is flipped: n rows, m cols
6    result := make([][]int, n)
7    for i := range result {
8        result[i] = make([]int, m)
9    }
10
11    for i := 0; i < n ; i++ {
12        for j := 0; j < m ; j++ {
13            result[i][j] = matrix[j][i]
14        }
15    }
16
17    return result
18}