1func matrixBlockSum(mat [][]int, k int) [][]int {
2    m := len(mat)
3    n := len(mat[0])
4
5    prefix := make([][]int, m+1)
6    for i := range prefix {
7        prefix[i] = make([]int, n+1)
8    }
9
10    for i := 1; i <= m; i++ {
11        for j := 1; j <= n; j++ {
12            prefix[i][j] = mat[i-1][j-1] +
13                        prefix[i-1][j] +
14                        prefix[i][j-1] - 
15                        prefix[i-1][j-1]
16        }
17    }
18
19    result := make([][]int, m)
20    for i := range result {
21        result[i] = make([]int, n)
22    }
23
24    for i := 0; i < m; i++ {
25        for j := 0; j < n; j++ {
26            r1 := max(0, i-k)
27			c1 := max(0, j-k)
28
29			r2 := min(m-1, i+k)
30			c2 := min(n-1, j+k)
31
32            r1++
33            r2++
34            c1++
35            c2++
36
37            result[i][j] = prefix[r2][c2] - 
38                        prefix[r1-1][c2] -
39                        prefix[r2][c1-1] +
40                        prefix[r1-1][c1-1]
41        }
42    }
43
44    return result
45}