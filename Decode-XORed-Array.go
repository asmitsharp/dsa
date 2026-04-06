1func decode(encoded []int, first int) []int {
2    res := make([]int, len(encoded)+1)
3    res[0] = first
4    for i := 0; i < len(encoded); i++ {
5        res[i+1] = res[i] ^ encoded[i]
6    }
7
8    return res
9}