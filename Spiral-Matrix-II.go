1func generateMatrix(n int) [][]int {
2    res := make([][]int, n)
3for i := range res {
4    res[i] = make([]int, n)
5}
6
7    top, bottom := 0, n-1
8    left, right := 0, n-1
9    num := 1
10
11    for top <= bottom && left <= right {
12        // travel right
13        for c := left; c <= right; c++ {
14            res[top][c] = num
15            num++
16        } 
17        top++
18
19        // travel down
20        for r := top; r <= bottom; r++ {
21            res[r][right] = num
22            num++
23        }
24        right--
25
26        // travel left
27        if top <= bottom {
28            for c := right; c >= left; c-- {
29                res[bottom][c] = num
30                num++
31            }
32            bottom--
33        }
34
35        // travel up 
36        if left <= right {
37            for r := bottom; r >= top; r-- {
38                res[r][left] = num
39                num++
40            }
41            left++
42        }
43    }
44
45    return res
46}