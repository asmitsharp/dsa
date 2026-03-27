1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func invertTree(root *TreeNode) *TreeNode {
10    if root == nil {
11        return root
12    }
13
14
15    root.Left, root.Right = root.Right, root.Left
16    invertTree(root.Left)
17    invertTree(root.Right)
18    return root
19}