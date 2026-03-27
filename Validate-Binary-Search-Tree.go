1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func isValidBST(root *TreeNode) bool {
10   var prev *TreeNode
11
12   var dfs func(node *TreeNode) bool
13   dfs = func(node *TreeNode) bool {
14    if node == nil {
15        return true
16    }
17
18    if !dfs(node.Left) {
19        return false
20    }
21
22    if prev != nil && node.Val <= prev.Val {
23        return false
24    }
25
26    prev = node
27
28    return dfs(node.Right)
29   }
30return dfs(root)
31}