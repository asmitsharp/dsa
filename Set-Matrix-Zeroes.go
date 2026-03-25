1func setZeroes(matrix [][]int)  {
2    m := len(matrix)
3    n := len(matrix[0])
4
5    firstRow := false
6    firstCol := false
7
8    for i := 0; i < m; i++ {
9        if matrix[i][0] == 0 {
10            firstCol = true
11        }
12    }
13
14    for j := 0; j < n; j++ {
15        if matrix[0][j] == 0 {
16            firstRow = true
17        }
18    }
19
20    for i := 1; i < m; i++ {
21        for j := 1; j < n; j++ {
22            if matrix[i][j] == 0 {
23                matrix[i][0] = 0
24                matrix[0][j] = 0
25            }
26        }
27    }
28
29    for i := 1; i < m; i++ {
30        for j := 1; j < n; j++ {
31            if matrix[i][0] == 0 || matrix[0][j] == 0  {
32                matrix[i][j] = 0
33            }
34        }
35    }
36
37    if firstCol {
38        for i := 0; i < m; i++ {
39            matrix[i][0] = 0
40        }
41    }
42
43    if firstRow {
44        for i := 0; i < n; i++ {
45            matrix[0][i] = 0
46        }
47    }
48}