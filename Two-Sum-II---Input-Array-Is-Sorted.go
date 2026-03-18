1func twoSum(numbers []int, target int) []int {
2    left := 0
3    right := len(numbers) - 1 
4
5    for left < right {
6        if numbers[left] + numbers[right] == target {
7            return []int{left+1, right+1}
8        } else if numbers[left] + numbers[right] < target {
9            left++
10        } else {
11            right--
12        }
13    }
14    return []int{}
15}