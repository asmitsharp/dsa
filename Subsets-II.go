1func subsetsWithDup(nums []int) [][]int {
2    res := [][]int{}
3    sort.Ints(nums)
4
5    backtrack(0, nums, []int{}, &res)
6    return res
7}
8
9func backtrack(start int, nums []int, curr []int, res *[][]int) {
10    temp := make([]int, len(curr))
11    copy(temp, curr)
12    *res = append(*res, temp)
13
14    for i := start; i < len(nums); i++ {
15        if i > start && nums[i] == nums[i-1] {
16            continue
17        }
18
19        curr = append(curr, nums[i])
20        backtrack(i+1, nums, curr, res)
21        curr = curr[:len(curr)-1]
22    }
23}