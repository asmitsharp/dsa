1func minimumXORSum(nums1 []int, nums2 []int) int {
2    n := len(nums1)
3    memo := make(map[int]int)
4
5    var dp func(mask int) int
6    dp = func(mask int) int {
7        if mask == (1<<n)-1 {
8            return 0
9        }
10
11        if v, ok := memo[mask]; ok {
12            return v
13        }
14
15        i := bits.OnesCount(uint(mask))
16        res := math.MaxInt32
17        for j := 0; j < n;  j++ {
18            if mask&(1<<j) == 0 {
19                res = min(res, (nums1[i]^nums2[j]) + dp(mask | (1<<j)))
20            }
21        }
22        memo[mask] = res
23        return res
24    }
25    return dp(0)
26}