1func xorOperation(n int, start int) int {
2    nums := make([]int, n)
3
4    for i , _ := range nums {
5        nums[i] = start + 2 * i
6    }
7
8    res := 0
9    for _, v := range nums {
10        res ^= v
11    }
12
13    return res
14}