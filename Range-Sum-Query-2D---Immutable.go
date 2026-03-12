1type NumMatrix struct {
2    prefix [][]int
3}
4
5
6func Constructor(matrix [][]int) NumMatrix {
7    
8    m := len(matrix)
9    n := len(matrix[0])
10
11    prefix := make([][]int, m+1)
12
13    for i := range prefix {
14        prefix[i] = make([]int, n+1)
15    }
16
17    for r := 1; r <= m; r++ {
18        for c := 1; c <= n; c++ {
19
20            prefix[r][c] =
21                matrix[r-1][c-1] +
22                prefix[r-1][c] +
23                prefix[r][c-1] -
24                prefix[r-1][c-1]
25        }
26    }
27
28    return NumMatrix{prefix: prefix}
29}
30
31
32func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
33    return this.prefix[row2+1][col2+1] - 
34        this.prefix[row1][col2+1] - 
35        this.prefix[row2+1][col1] + 
36        this.prefix[row1][col1]
37}
38
39
40/**
41 * Your NumMatrix object will be instantiated and called as such:
42 * obj := Constructor(matrix);
43 * param_1 := obj.SumRegion(row1,col1,row2,col2);
44 */