1func findDuplicate(nums []int) int {
2    i := 0
3    n := len(nums)
4
5    for i < n {
6        correctIdx := nums[i] - 1
7        if nums[i] < n && nums[i] != i+1 {
8            if nums[i] != nums[correctIdx] {
9                nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
10            } else {
11                return nums[i]
12            }
13        } else {
14            i++
15        }
16    }
17    return -1
18}