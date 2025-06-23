import "strings"

func isValidSerialization(preorder string) bool {
    if preorder == "" {
        return false
    }

    nodes := strings.Split(preorder, ",")

    stack := make([]string, 0)

    for _, node := range nodes {
        stack = append(stack, node)

        for len(stack) >= 3 && 
                stack[len(stack) - 1] == "#" && 
                stack[len(stack) - 2] == "#" && 
                stack[len(stack) - 3] != "#" {
            stack = stack[:len(stack) - 3]
            stack = append(stack, "#")
        } 
    }

    return len(stack) == 1 && stack[0] == "#" 
}