1func carPooling(trips [][]int, capacity int) bool {
2    diff := make([]int, 1001)
3
4    for _, trip := range trips {
5        passenger := trip[0]
6        from := trip[1]
7        to := trip[2]
8
9        diff[from] += passenger
10        diff[to] -= passenger
11    }
12
13    current := 0
14
15    for i := 0; i < 1001; i++ {
16        current += diff[i]
17
18        if current > capacity {
19            return false
20        }
21    }
22
23    return true
24}