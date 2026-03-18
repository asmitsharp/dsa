1func numberOfSubarrays(nums []int, k int) int {
2    return atmost(nums, k) - atmost(nums, k-1)
3}
4
5func atmost(nums []int, k int) int {
6    left := 0
7    count := 0
8    oddCount := 0
9
10    for right := 0; right < len(nums); right++ {
11       if nums[right] % 2 != 0 {
12            oddCount++
13        }
14
15        for oddCount > k {
16            if nums[left] % 2 != 0 {
17                oddCount--
18            }
19            left++
20        } 
21
22        count += right - left + 1
23    }
24    return count
25}