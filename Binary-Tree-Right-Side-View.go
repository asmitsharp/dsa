1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func rightSideView(root *TreeNode) []int {
10    res := []int{}
11
12    var dfs func(node *TreeNode, depth int)
13    dfs = func(node *TreeNode, depth int) {
14        if node == nil {
15            return
16        }
17
18        if depth == len(res) {
19            res = append(res, node.Val)
20        }
21
22        dfs(node.Right, depth+1)
23        dfs(node.Left, depth+1)
24    }
25
26    dfs(root, 0)
27    return res
28}