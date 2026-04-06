1func findTheDifference(s string, t string) byte {
2    var res byte = 0
3
4    for i := 0; i < len(s); i++ {
5        res ^= s[i]
6    }
7
8    for i := 0; i < len(t); i++ {
9        res ^= t[i]
10    }
11
12    return res
13}