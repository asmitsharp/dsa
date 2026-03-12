1func findMiddleIndex(nums []int) int {
2    total := 0
3    for _, v := range nums {
4        total += v
5    }
6
7    leftSum := 0
8    for i := 0; i < len(nums); i++ {
9        rightSum := total - leftSum - nums[i]
10
11        if leftSum == rightSum {
12            return i
13        }
14        leftSum += nums[i]
15    }
16    return -1
17}