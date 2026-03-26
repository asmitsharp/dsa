1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func isSymmetric(root *TreeNode) bool {
10    return isMirror(root.Left, root.Right)
11}
12
13func isMirror(l, r *TreeNode) bool {
14    if l == nil && r == nil {return true}
15    if l == nil || r == nil {return false}
16
17    return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
18}