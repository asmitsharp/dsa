1func numSubarrayProductLessThanK(nums []int, k int) int {
2    if k <= 1 {
3        return 0
4    }
5    left := 0
6    winPro := 1
7    count := 0
8
9    for right := 0; right < len(nums); right++ {
10        winPro *= nums[right]
11
12        for winPro >= k {
13            winPro /= nums[left]
14            left++
15        }
16
17        count += right - left + 1
18    }
19    return count
20}