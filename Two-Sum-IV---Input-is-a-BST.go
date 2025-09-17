/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func findTarget(root *TreeNode, k int) bool {
//     arr := inorder(root)

//     l, r := 0, len(arr) - 1
//     for l < r {
//         sum := arr[l] + arr[r]
//         if sum == k {
//             return true  
//         } else if sum < k {
//             l++
//         } else {
//             r--
//         }
//     }
//     return false
// }

// func inorder(node *TreeNode) []int {
//     if node == nil {
//         return nil
//     }

//     left := inorder(node.Left)
//     curr := []int{node.Val}
//     right := inorder(node.Right)

//     return append(append(left, curr...), right...)
// }

func findTarget(root *TreeNode, k int) bool {
    seen := make(map[int]bool)
    return dfs(root, k, seen)
}

func dfs(node *TreeNode, k int, seen map[int]bool) bool {
    if node == nil {
        return false
    }

    if seen[k - node.Val] {
        return true
    }

    seen[node.Val] = true

    return dfs(node.Left, k, seen) || dfs(node.Right, k, seen)
}