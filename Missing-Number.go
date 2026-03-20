1func missingNumber(nums []int) int {
2    i := 0
3    for i < len(nums) {
4        correctIdx := nums[i]
5        if nums[i] < len(nums) && nums[i] != nums[correctIdx] {
6            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
7        } else {
8            i++
9        }
10    }
11
12    for i := 0; i < len(nums); i++ {
13        if nums[i] != i {
14            return i
15        }
16    }
17
18    return len(nums)
19}