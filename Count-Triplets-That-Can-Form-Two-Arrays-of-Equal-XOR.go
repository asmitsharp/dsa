1func countTriplets(arr []int) int {
2    count := 0
3
4    for i := 0; i < len(arr); i++ {
5        xor := 0
6        for k := i; k < len(arr); k++ {
7            xor ^= arr[k]
8
9            if xor == 0 {
10                count += k-i
11            }
12        }
13    }
14
15    return count
16}