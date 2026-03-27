1func combinationSum(candidates []int, target int) [][]int {
2    result := [][]int{}
3    backtrack(0,target, []int{}, &result, candidates)
4    return result
5}
6
7func backtrack(start int, target int, curr []int, result *[][]int, candidates []int) {
8    if target == 0 {
9        temp := make([]int, len(curr))
10        copy(temp, curr)
11        *result = append(*result, temp)
12    }
13
14    if target < 0 {
15        return
16    }
17
18    for i := start; i < len(candidates); i++ {
19        curr = append(curr, candidates[i])
20        backtrack(i, target-candidates[i],curr, result, candidates)
21        curr = curr[:len(curr)-1]
22    }
23}