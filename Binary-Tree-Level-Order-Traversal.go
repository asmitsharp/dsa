1/**
2 * Definition for a binary tree node.
3 * type TreeNode struct {
4 *     Val int
5 *     Left *TreeNode
6 *     Right *TreeNode
7 * }
8 */
9func levelOrder(root *TreeNode) [][]int {
10    if root == nil {
11        return nil
12    }
13
14    result := [][]int{}
15    queue := []*TreeNode{root}
16
17    for len(queue) > 0 {
18        size := len(queue) 
19        level := []int{}
20
21        for i := 0; i < size; i++ {
22            node := queue[0]
23            queue = queue[1:]
24
25            level = append(level, node.Val)
26
27            if node.Left != nil {
28                queue = append(queue, node.Left)
29            }
30            if node.Right != nil {
31                queue = append(queue, node.Right)
32            }
33
34        }
35        result = append(result, level)
36    }
37    return result
38}