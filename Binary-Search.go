1func search(nums []int, target int) int {
2    left, right := 0, len(nums)-1
3
4    for left <= right {
5        mid := left + (right-left) / 2
6
7        if nums[mid] == target {
8            return mid
9        } else if nums[mid] < target {
10            left = mid + 1
11        } else {
12            right = mid - 1
13        }
14    }
15    return -1
16}