func minReorder(n int, connections [][]int) int {
    forward_nbr := make(map[int][]int)
    backward_nbr := make(map[int][]int)

    visited := make(map[int]bool)

    for _, conn := range connections {
        a := conn[0]
        b := conn[1]

        forward_nbr[a] = append(forward_nbr[a], b)
        backward_nbr[b] = append(backward_nbr[b], a)
    }

    ans := 0
    dfs(0, forward_nbr, backward_nbr, &ans, visited)
    return ans
}

func dfs(source int, forward_nbr, backward_nbr map[int][]int, ans *int, visited map[int]bool) {
    visited[source] = true

    for _, nbr := range forward_nbr[source] {
        if !visited[nbr] {
            *ans+=1
            dfs(nbr, forward_nbr, backward_nbr, ans, visited)
        }
    }

    for _, nbr := range backward_nbr[source] {
        if !visited[nbr] {
            dfs(nbr, forward_nbr, backward_nbr, ans, visited)
        }
    }
}