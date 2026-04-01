1// func searchInsert(nums []int, target int) int {
2//     left, right := 0, len(nums)
3//     for left < right {
4//         mid := left + (right - left) / 2
5//         if target > nums[mid] {
6//             left = mid + 1
7//         } else {
8//             right = mid
9//         } 
10//     }
11//     return left
12// }
13
14
15func searchInsert(nums []int, target int) int {
16   i := 0
17
18   for i < len(nums) && nums[i] < target {
19    i++
20   } 
21
22   return i
23}
24
25