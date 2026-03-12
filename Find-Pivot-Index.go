1func pivotIndex(nums []int) int {
2    total := 0
3    for _, v := range nums {
4        total += v
5    }
6
7    leftSum := 0
8    for i := 0; i < len(nums); i++ {
9        rightSum := total - leftSum - nums[i]
10        if leftSum == rightSum {
11            return i
12        }
13        leftSum += nums[i]
14    }
15
16    return -1
17}