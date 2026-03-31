1func solveNQueens(n int) [][]string {
2    board := make([][]byte, n)
3    for i := range board {
4        board[i] = make([]byte, n)
5        for j := range board[i] {
6            board[i][j] = '.'
7        }
8    }
9    result := [][]string{}
10
11    var backtrack func(start int)
12    backtrack = func(start int) {
13        if start == n {
14            temp := make([]string, n)
15            for i := 0; i < n; i++ {
16                temp[i] = string(board[i])
17            }
18            result = append(result, temp)
19            return
20        }
21
22        for col := 0; col < n;  col++ {
23            if isSafe(board, start, col, n) {
24                board[start][col] = 'Q'
25                backtrack(start+1)
26                board[start][col] = '.'
27            }
28        }
29    }
30    backtrack(0)
31    return result
32}
33
34func isSafe(board [][]byte, row, col int, n int) bool {
35    for i := 0; i < row; i++ {
36        if board[i][col] == 'Q' {
37            return false
38        }
39    }
40
41     for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
42        if board[i][j] == 'Q' {
43            return false
44        }
45    }
46
47     for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
48        if board[i][j] == 'Q' {
49            return false
50        }
51    }
52
53    return true
54}