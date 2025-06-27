import "strings"

func lengthLongestPath(input string) int {
    lines := strings.Split(input, "\n")
    pathLen := map[int]int{}
    maxLen := 0

    for _, line := range lines {
        depth := strings.Count(line, "\t")
        name := strings.TrimLeft(line, "\t")

        if strings.Contains(name, ".") {
            totalLen := pathLen[depth-1] + len(name)
            if totalLen > maxLen {
                maxLen = totalLen
            }
        } else {
            pathLen[depth] = pathLen[depth-1] + len(name) + 1
        }

    }
    return maxLen
}