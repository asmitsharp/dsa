1func missingNumber(nums []int) int {
2    res := len(nums)
3
4    for i, n := range nums {
5        res ^= i ^ n
6    }
7
8    return res
9}