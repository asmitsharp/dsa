1func matrixReshape(mat [][]int, r int, c int) [][]int {
2    if len(mat) * len(mat[0]) != r * c {
3        return mat
4    }
5
6    res := make([][]int, r)
7    for i := range res {
8        res[i] = make([]int, c)
9    }
10    m := len(mat[0])
11
12    for k := 0; k < r*c; k++ {
13        res[k/c][k%c] = mat[k/m][k%m]
14    }
15
16    return res
17
18}