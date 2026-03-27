1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func pathSum(root *TreeNode, targetSum int) [][]int {
10    result := [][]int{}
11    path := []int{}
12    
13    var dfs func(node *TreeNode, sum int)
14    dfs = func(node *TreeNode, sum int) {
15        if node == nil {
16            return
17        }
18
19        path = append(path, node.Val)
20
21        if node.Left == nil && node.Right == nil && sum == node.Val {
22            temp := make([]int, len(path))
23            copy(temp, path)
24            result = append(result, temp)
25        }
26
27        dfs(node.Left, sum-node.Val)
28        dfs(node.Right, sum-node.Val)
29
30        path = path[:len(path)-1]
31    }
32
33    dfs(root, targetSum)
34    return result
35}
36
37