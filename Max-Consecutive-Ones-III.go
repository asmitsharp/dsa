1func longestOnes(nums []int, k int) int {
2    left := 0
3    maxLen := 0
4    zeroCount := 0
5
6    for right := 0; right < len(nums); right++ {
7        if nums[right] == 0 {
8            zeroCount++
9        }
10
11        for zeroCount > k {
12            if nums[left] == 0 {
13                zeroCount--
14            }
15            left++
16        }
17
18        maxLen = max(maxLen, right - left + 1)
19    }
20
21    return maxLen
22}
23
24func max(a, b int) int {
25    if a > b {
26        return a
27    }
28    return b
29}