1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func isSameTree(p *TreeNode, q *TreeNode) bool {
10    if p == nil && q == nil {return true}
11    if p == nil || q == nil {return false}
12
13    return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
14}