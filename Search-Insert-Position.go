// func searchInsert(nums []int, target int) int {
//     left, right := 0, len(nums)
//     for left < right {
//         mid := left + (right - left) / 2
//         if target > nums[mid] {
//             left = mid + 1
//         } else {
//             right = mid
//         } 
//     }
//     return left
// }


func searchInsert(nums []int, target int) int {
   i := 0

   for i < len(nums) && nums[i] < target {
    i++
   } 

   return i
}

