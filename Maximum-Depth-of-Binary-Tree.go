1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func maxDepth(root *TreeNode) int {
10    if root == nil {
11        return 0
12    }
13
14    left := maxDepth(root.Left)
15    right := maxDepth(root.Right)
16
17    return 1 + max(left, right)
18}
19
20func max(a, b int) int {
21    if a > b {
22        return a
23    }
24    return b
25}