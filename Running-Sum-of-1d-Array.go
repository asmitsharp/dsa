1func runningSum(nums []int) []int {
2    result := make([]int, len(nums))
3
4    result[0] = nums[0]
5    for i := 1; i < len(nums); i++ {
6        result[i] = result[i-1] + nums[i]
7    }
8
9    return result
10}