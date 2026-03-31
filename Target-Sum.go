1func findTargetSumWays(nums []int, target int) int {
2    count := 0
3
4    var backtrack func(start int, sum int)
5    backtrack = func(start int, sum int) {
6        if start == len(nums) {
7            if sum == target {
8                count++
9            }
10            return
11        }
12
13        backtrack(start+1,sum+nums[start])
14        backtrack(start+1,sum-nums[start])
15    }
16    backtrack(0,0)
17    return count
18}