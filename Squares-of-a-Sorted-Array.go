func sortedSquares(nums []int) []int {
   n := len(nums)
   result := make([]int, n)

   l := 0
   r := n - 1
   p := n - 1

   for l <= r {
    left := nums[l] * nums[l]
    right := nums[r] * nums[r]

    if left > right {
        result[p] = left
        l++
    } else {
        result[p] = right
        r--
    }
    p--
   }

   return result
}