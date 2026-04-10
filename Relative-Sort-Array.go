1func relativeSortArray(arr1 []int, arr2 []int) []int {
2    count := make(map[int]int)
3
4    for _ , v := range arr1 {
5        count[v]++
6    }
7
8    res := []int{}
9
10    for _ , v := range arr2 {
11        for count[v] > 0 {
12            res = append(res, v)
13            count[v]--
14        }
15    }
16
17    rest := []int{}
18    for k, v := range count {
19        for v > 0 {
20            rest = append(rest, k)
21            v--
22        }
23    }
24
25    sort.Ints(rest)
26    res = append(res, rest...)
27    return res
28}