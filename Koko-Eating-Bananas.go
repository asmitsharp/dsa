1func minEatingSpeed(piles []int, h int) int {
2    lo,hi := 1,0
3    for _ , p := range piles {
4        if hi < p {
5            hi = p
6        }
7    }
8
9    for lo < hi {
10        mid := lo + (hi - lo) / 2
11
12        if canFinish(piles, mid, h) {
13            hi = mid
14        } else {
15            lo = mid + 1
16        }
17    }
18    return lo
19}
20
21func canFinish(piles []int, speed , h int) bool {
22    hours := 0 
23    for _, p := range piles {
24        hours += (p + speed - 1) / speed
25    }
26    return hours <= h
27}