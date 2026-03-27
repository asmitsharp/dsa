1func permute(nums []int) [][]int {
2    result := [][]int{}
3    used := make([]bool, len(nums))
4    backtrack([]int{}, &result, nums, used)
5    return result
6}
7
8func backtrack(curr []int, result *[][]int, nums []int, used []bool) {
9    if len(curr) == len(nums) {
10        temp := make([]int, len(curr))
11        copy(temp, curr)
12        *result = append(*result, temp)
13        return
14    }
15
16    for i := 0; i < len(nums); i++ {
17        if used[i] {
18            continue
19        }
20        used[i] = true
21        curr = append(curr, nums[i])
22        backtrack(curr, result, nums, used)
23        curr = curr[:len(curr)-1]
24        used[i] = false
25    }
26}