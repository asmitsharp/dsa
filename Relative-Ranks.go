1func findRelativeRanks(score []int) []string {
2    n := len(score)
3    res := make([]string, n)
4
5    type pair struct {
6        val int
7        idx int
8    }
9
10    arr := make([]pair, n)
11
12    for i, val := range score {
13        arr[i] = pair{val, i}
14    }
15
16    sort.Slice(arr, func(i, j int) bool {
17        return arr[i].val > arr[j].val
18    })
19
20    for i , p := range arr {
21        switch i {
22            case 0:
23            res[p.idx] = "Gold Medal"
24            case 1:
25            res[p.idx] = "Silver Medal"
26            case 2:
27            res[p.idx] = "Bronze Medal"
28            default :
29            res[p.idx] = strconv.Itoa(i+1)
30        }
31    }
32    return res
33}