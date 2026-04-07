1func getXORSum(arr1 []int, arr2 []int) int {
2    xor1 , xor2 := 0, 0
3    for i := 0; i < len(arr1); i++ {
4        xor1 ^= arr1[i]
5    }
6
7    for j := 0; j < len(arr2); j++ {
8        xor2 ^= arr2[j]
9    }
10
11    return xor1 & xor2
12}