1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func maxPathSum(root *TreeNode) int {
10    maxSum := math.MinInt32
11
12    var dfs func(node *TreeNode) int 
13    dfs = func(node *TreeNode) int {
14        if node == nil {
15            return 0
16        }
17
18        left := max(0, dfs(node.Left))
19        right := max(0, dfs(node.Right))
20
21        maxSum = max(maxSum, left+right+node.Val)
22        return node.Val+max(left, right)
23    }
24
25    dfs(root)
26    return maxSum
27}