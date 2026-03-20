1func findDisappearedNumbers(nums []int) []int {
2    i := 0
3    n := len(nums)
4    for i < n {
5        correctIdx := nums[i] - 1
6        if nums[i] != nums[correctIdx] {
7            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
8        } else {
9            i++
10        }
11    }
12
13    result := []int{}
14    for i := 0; i < n; i++ {
15        if nums[i] != i + 1 {
16            result = append(result, i+1)
17        }
18    }
19
20    return result
21}