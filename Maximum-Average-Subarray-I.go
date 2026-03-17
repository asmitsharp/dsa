1func findMaxAverage(nums []int, k int) float64 {
2  
3    
4    windowSum := 0
5    for i := 0; i < k; i++ {
6        windowSum += nums[i]
7    }
8
9    maxSum := windowSum
10
11    for i := k; i < len(nums); i++ {
12        windowSum = windowSum - nums[i-k] + nums[i]
13
14        if windowSum > maxSum {
15            maxSum = windowSum
16        }
17    }
18
19    return float64(maxSum) / float64(k)
20}