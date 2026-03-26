1func updateMatrix(mat [][]int) [][]int {
2    row := len(mat)
3    col := len(mat[0])
4
5    queue := [][]int{}
6
7    for r := 0; r < row; r++ {
8        for c := 0; c < col; c++ {
9            if mat[r][c] == 0 {
10                queue = append(queue, []int{r,c})
11            } else {
12                mat[r][c] = -1
13            }
14        }
15    }
16
17    dirs := [][]int{
18        {1,0},
19        {-1,0},
20        {0,1},
21        {0,-1},
22    }
23
24    for len(queue) > 0 {
25        cell := queue[0]
26        queue = queue[1:]
27
28        r, c := cell[0], cell[1]
29
30        for _, d := range dirs {
31            nr := r + d[0]
32            nc := c + d[1]
33
34            if nr >= 0 && nr < row && nc >= 0 && nc < col && mat[nr][nc] == -1 {
35                mat[nr][nc] = mat[r][c]+1
36                queue = append(queue, []int{nr, nc})
37            }
38        }
39    }
40
41    return mat
42}