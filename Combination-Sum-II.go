1func combinationSum2(candidates []int, target int) [][]int {
2    res := [][]int{}
3    sort.Ints(candidates)
4    backtrack(0, candidates, []int{}, target, &res)
5    return res
6}
7
8func backtrack(start int, cand []int, curr []int, target int, res *[][]int) {
9    if target == 0 {
10        temp := make([]int, len(curr))
11        copy(temp, curr)
12        *res = append(*res, temp)
13        return
14    }
15
16    if target < 0 {
17        return
18    }
19
20    for i := start; i < len(cand); i++ {
21        if i > start && cand[i] == cand[i-1] {
22            continue
23        }
24        curr = append(curr, cand[i])
25        backtrack(i+1, cand, curr, target-cand[i], res)
26        curr = curr[:len(curr)-1]
27    }
28}