1func rotate(matrix [][]int)  {
2    n := len(matrix)
3
4    for i := 0; i < n; i++ {
5        for j := i+1; j < n; j++ {
6            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
7        }
8    }
9
10    for i := 0; i < n; i++ {
11        l, r := 0, n-1
12        for  l < r {
13            matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
14            l++
15            r--
16        }
17    }
18}