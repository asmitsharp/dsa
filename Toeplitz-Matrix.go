1func isToeplitzMatrix(matrix [][]int) bool {
2   for i := 1; i < len(matrix); i++ {
3    for j := 1; j < len(matrix[0]); j++ {
4        if matrix[i][j] != matrix[i-1][j-1] {
5            return false
6        }
7    }
8   }
9   return true
10}