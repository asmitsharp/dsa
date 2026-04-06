1func singleNumber(nums []int) []int {
2    xor := 0
3    for _ , n := range nums {
4        xor ^= n
5    }
6
7    diffBit := xor & (-xor)
8
9    x,y := 0,0
10    for _, n := range nums {
11        if n & diffBit != 0 {
12            x ^= n
13        } else {
14            y ^= n
15        }
16    }
17
18     return []int{x, y}
19}