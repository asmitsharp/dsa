1func findMaxLength(nums []int) int {
2    prefix := 0
3    maxLen := 0
4
5    firstIndex := map[int]int {
6        0 : -1,
7    }
8
9    for i, v := range nums {
10
11        if v == 0 {
12            prefix -= 1
13        } else {
14            prefix += 1
15        }
16
17        if idx, ok := firstIndex[prefix]; ok {
18            length  := i - idx
19            if length > maxLen {
20                maxLen = length
21            }
22        } else {
23            firstIndex[prefix] = i
24        }
25    }
26
27    return maxLen
28}