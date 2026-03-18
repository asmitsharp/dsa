1func totalFruit(fruits []int) int {
2    // 2 baskets -> each basket ( single type of fruit )
3    // one fruit from everyb tree
4    left := 0
5    freq := make(map[int]int)
6    maxFruits := 0
7
8    for right := 0; right < len(fruits); right++ {
9        freq[fruits[right]]++
10
11        for len(freq) > 2 {
12            freq[fruits[left]]--
13
14            if freq[fruits[left]] == 0 {
15                delete(freq, fruits[left])
16            }
17
18            left++ 
19        }
20        maxFruits = max(maxFruits, right-left+1)
21    }
22    return maxFruits
23}
24
25func max(a, b int) int {
26    if a > b {
27        return a
28    } 
29    return b
30}