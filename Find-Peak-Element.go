1func findPeakElement(nums []int) int {
2    left , right := 0 , len(nums) - 1
3
4    for left < right {
5        mid := left + (right - left) / 2
6
7        if nums[mid] < nums[mid+1] {
8            left = mid + 1
9        } else {
10            right = mid
11        }
12    }
13    return left
14}