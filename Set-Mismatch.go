1func findErrorNums(nums []int) []int {
2    i := 0
3    n := len(nums) 
4
5    for i < n {
6        correctIdx := nums[i] - 1
7        if nums[i] != nums[correctIdx] {
8            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
9        } else {
10            i++
11        }
12    }
13
14    for i := 0; i < n; i++ {
15        if nums[i] != i+1 {
16            return []int{nums[i], i+1}
17        }
18    }
19    return []int{-1, -1}
20}