1func spiralOrder(matrix [][]int) []int {
2    res := []int{}
3    top , bottom := 0, len(matrix) - 1
4    left , right := 0, len(matrix[0]) - 1
5
6    for top <= bottom && left <= right {
7        // travel right
8        for c := left; c <= right; c++ {
9            res = append(res, matrix[top][c])
10        }
11        top++
12
13        // travel down
14        for r := top; r <= bottom; r++ {
15            res = append(res, matrix[r][right])
16        }
17        right--
18
19        // travel left
20        if top <= bottom {
21            for c := right; c >= left; c-- {
22                res = append(res, matrix[bottom][c])
23            }
24            bottom--
25        }
26        // travel up
27        if left <= right {
28            for r := bottom; r >= top; r-- {
29                res = append(res, matrix[r][left])
30            }
31            left++
32        }
33    }
34    return res
35}