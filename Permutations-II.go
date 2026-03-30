1func permuteUnique(nums []int) [][]int {
2    res := [][]int{}
3    sort.Ints(nums)
4    used := make([]bool, len(nums))
5    backtrack([]int{}, nums, used, &res)
6    return res
7}
8
9func backtrack(curr []int, nums []int, used []bool, res *[][]int) {
10    if len(curr) == len(nums) {
11        temp := make([]int, len(curr))
12        copy(temp, curr)
13        *res = append(*res, temp)
14        return
15    }
16
17    for i := 0; i < len(nums); i++ {
18        if used[i]  {
19            continue
20        }
21
22        if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
23            continue
24        }
25        used[i] = true
26        curr = append(curr, nums[i])
27        backtrack(curr, nums, used, res)
28        curr = curr[:len(curr)-1]
29        used[i] = false
30    }
31}