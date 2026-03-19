1func sortedSquares(nums []int) []int {
2    result := make([]int, len(nums))
3
4    left := 0
5    right := len(nums) - 1
6    pos := len(nums) - 1
7
8    for left <= right {
9        squareLeft := nums[left] * nums[left]
10        squareRight := nums[right] * nums[right]
11
12        if squareLeft < squareRight {
13            result[pos] = squareRight
14            right--
15        } else {
16            result[pos] = squareLeft
17            left++
18        }
19        pos--
20    }
21
22    return result
23}