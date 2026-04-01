1func search(nums []int, target int) bool {
2    left , right := 0, len(nums) - 1
3
4    for left <= right {
5
6        mid := left + (right - left) / 2
7
8        if nums[mid] == target {
9            return true
10        }
11
12        if nums[left] == nums[mid] && nums[mid] == nums[right] {
13            left++
14            right--
15            continue
16        }
17
18        if nums[left] <= nums[mid] {
19            if nums[left] <= target && target < nums[mid] {
20                right = mid - 1
21            } else {
22                left = mid + 1
23            }
24        } else {
25            if nums[right] >= target && nums[mid] < target {
26                left = mid + 1
27            } else {
28                right = mid - 1
29            }
30        }
31    }
32    return false
33}