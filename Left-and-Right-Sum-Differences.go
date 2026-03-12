1func leftRightDifference(nums []int) []int {
2    total := 0
3    for _,v := range nums {
4        total += v
5    }
6
7    answer := make([]int, len(nums))
8    leftSum := 0
9    for i := 0; i < len(nums); i++ {
10        rightSum := total - leftSum - nums[i]
11
12        diff := leftSum - rightSum
13        if diff < 0 {
14            diff = -diff
15        }
16        answer[i] = diff
17        leftSum += nums[i]
18    }
19
20    return answer
21}