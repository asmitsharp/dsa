1func xorQueries(arr []int, queries [][]int) []int {
2    n := len(arr)
3
4    prefix := make([]int, n)
5    prefix[0] = arr[0]
6
7    for i := 1; i < n; i++ {
8        prefix[i] = prefix[i-1] ^ arr[i]
9    }
10
11    result := make([]int, len(queries))
12
13    for i, q := range queries {
14        left := q[0]
15        right := q[1] 
16
17        if left == 0 {
18            result[i] = prefix[right]
19        } else {
20            result[i] = prefix[right] ^ prefix[left - 1]
21        }
22    }
23
24    return result
25}