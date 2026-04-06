1func singleNumber(nums []int) int {
2    result := 0
3
4    for _, n := range nums {
5        result ^= n
6    }
7
8    return result
9}