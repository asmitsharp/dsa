1func findDuplicates(nums []int) []int {
2    i := 0
3    n := len(nums)
4
5    for i < n {
6        correctIdx := nums[i] - 1
7        if nums[i] != nums[correctIdx] {
8            nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
9        } else {
10            i++
11        }
12    }
13
14    result := []int{}
15    for i, num := range nums {
16        if num != i+1 {
17            result = append(result, num)
18        }
19    }
20
21    return result
22}