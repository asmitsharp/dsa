1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func goodNodes(root *TreeNode) int {
10    var dfs func(node *TreeNode, maxVal int) int 
11    dfs = func(node *TreeNode, maxVal int) int {
12        if node == nil {
13        return 0
14        }
15        count := 0
16        if node.Val >= maxVal {
17            count = 1
18            maxVal = node.Val
19        }
20
21        count += dfs(node.Left, maxVal)
22        count += dfs(node.Right, maxVal)
23
24        return count
25    }
26    return dfs(root, root.Val)
27}