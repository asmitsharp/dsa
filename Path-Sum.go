1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func hasPathSum(root *TreeNode, targetSum int) bool {
10    if root == nil {
11        return false
12    }
13
14    if root.Left == nil && root.Right == nil {
15        return root.Val == targetSum
16    }
17
18    remaining := targetSum - root.Val
19
20    return hasPathSum(root.Left, remaining) || hasPathSum(root.Right, remaining)
21}