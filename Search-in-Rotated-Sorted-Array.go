1func search(nums []int, target int) int {
2    left , right := 0, len(nums) - 1
3
4    for left <= right {
5        mid := left + (right - left) / 2
6
7        if nums[mid] == target {
8            return mid
9        }
10
11        if nums[left] <= nums[mid] {
12            if nums[left] <= target && target < nums[mid] {
13                right = mid - 1
14            } else {
15                left = mid + 1
16            }
17        } else {
18            if nums[right] >= target && target > nums[mid] {
19                left = mid + 1
20            } else {
21                right = mid - 1
22            }
23        }
24    }
25    return -1
26}