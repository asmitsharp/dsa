1func minSubArrayLen(target int, nums []int) int {
2    left := 0
3    minLen := len(nums) + 1
4    winSum := 0
5
6    for right := 0; right < len(nums); right++ {
7        winSum += nums[right]
8
9        for winSum >= target {
10            winLen := right - left + 1
11            if  winLen < minLen {
12                minLen = winLen
13            }
14
15            winSum -= nums[left]
16            left++
17        }
18    }
19
20    if minLen == len(nums) + 1 {
21        return 0
22    }
23    return minLen
24}