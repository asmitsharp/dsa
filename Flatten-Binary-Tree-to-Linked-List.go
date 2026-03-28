1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func flatten(root *TreeNode)  {
10    var prev *TreeNode
11
12    var dfs func(node *TreeNode) 
13    dfs = func(node *TreeNode) {
14        if node == nil {
15            return
16        }
17
18        dfs(node.Right)
19        dfs(node.Left)
20
21        node.Right = prev
22        node.Left = nil
23        prev = node
24    }
25
26    dfs(root)
27}