1func combine(n int, k int) [][]int {
2    result := [][]int{}
3    path := []int{}
4
5    var backtrack func(start int)
6    backtrack = func(start int) {
7        if len(path) == k {
8            temp := make([]int, len(path))
9            copy(temp, path)
10            result = append(result, temp)
11            return
12        }
13
14        for i := start; i <= n; i++ {
15            path = append(path, i)
16            backtrack(i+1)
17            path = path[:len(path)-1]
18        }
19    }
20    backtrack(1)
21    return result
22}