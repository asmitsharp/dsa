1func removeDuplicates(nums []int) int {
2    if len(nums) == 0 {
3        return 0
4    }
5
6    slow := 0
7    for fast := 1; fast < len(nums); fast++ {
8        if nums[slow] != nums[fast] {
9            slow++
10            nums[slow] = nums[fast]
11        }
12    }
13    return slow + 1
14}