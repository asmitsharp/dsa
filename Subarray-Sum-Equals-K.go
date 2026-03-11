1func subarraySum(nums []int, k int) int {
2      count := 0
3    prefixSum := 0
4    freq := map[int]int{0: 1}
5
6    for _, num := range nums {
7        prefixSum += num
8        count += freq[prefixSum-k]
9        freq[prefixSum]++
10    }
11    return count
12}