1func subarraysDivByK(nums []int, k int) int {
2
3    remainderCount := make(map[int]int)
4    remainderCount[0] = 1
5
6    prefixSum := 0
7    result := 0
8
9    for _, num := range nums {
10        prefixSum += num
11
12        remainder := prefixSum % k
13        if remainder < 0 {
14            remainder += k
15        }
16
17        if count, exists := remainderCount[remainder]; exists {
18            result += count
19        }
20
21        remainderCount[remainder]++
22    }
23    return result
24}
25
26
27    
28