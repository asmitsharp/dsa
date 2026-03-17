1func numOfSubarrays(arr []int, k int, threshold int) int {
2    winSum := 0
3    count := 0
4    for i := 0; i < k; i++ {
5        winSum += arr[i]
6    }
7    if winSum / k >= threshold {
8        count++
9    }
10
11    for i := k; i < len(arr); i++ {
12        winSum += arr[i]
13        winSum -= arr[i-k]
14
15        if winSum / k >= threshold {
16            count++
17        }
18    }
19
20    return count
21}