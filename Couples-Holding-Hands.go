1func minSwapsCouples(row []int) int {
2    n := len(row)
3    pos := make([]int, n)
4
5    for i,person := range row {
6        pos[person] = i
7    }
8
9    swaps := 0
10
11    for i := 0; i < n; i+=2 {
12        x := row[i]
13        partner := 0
14
15        if x % 2 == 0 {
16            partner = x + 1
17        } else {
18            partner = x - 1
19        }
20
21        if row[i+1] != partner {
22            swaps++
23            correctIdx := pos[partner]
24
25            row[i+1], row[correctIdx] = row[correctIdx], row[i+1]
26            
27            pos[row[correctIdx]] = correctIdx
28            pos[row[i+1]] = i+1
29        }
30    }
31
32    return swaps
33}