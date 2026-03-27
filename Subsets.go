1func subsets(nums []int) [][]int {
2    result := [][]int{}
3    backtrack(0, []int{}, &result, nums)
4    return result
5}
6
7func backtrack(start int, curr []int, result *[][]int, nums []int)  {
8    temp := make([]int, len(curr))
9    copy(temp, curr)
10    *result = append(*result, temp)
11
12    for i := start; i < len(nums); i++ {
13        curr = append(curr, nums[i])
14        backtrack(i+1, curr, result, nums)
15        curr = curr[:len(curr) - 1]
16    }
17}