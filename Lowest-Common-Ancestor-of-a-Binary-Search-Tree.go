1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val   int
5 *     Left  *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9
10func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
11	if root == nil {
12        return root
13    }
14
15    if p.Val < root.Val && q.Val < root.Val {
16      return  lowestCommonAncestor(root.Left, p, q)
17    } else if p.Val > root.Val && q.Val > root.Val {
18       return lowestCommonAncestor(root.Right, p, q)
19    }
20    return root
21}