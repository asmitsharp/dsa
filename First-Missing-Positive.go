1func firstMissingPositive(nums []int) int {
2    i := 0
3    n := len(nums)
4
5    for i < n {
6        if nums[i] > 0 && nums[i] <= n {
7            correctIdx := nums[i] - 1
8            if nums[i] != nums[correctIdx] {
9            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
10            continue
11            } 
12        }
13        i++
14    }
15
16    for i := 0; i < n; i++ {
17        if nums[i] != i+1 {
18            return i+1
19        }
20    }
21    return n+1
22}