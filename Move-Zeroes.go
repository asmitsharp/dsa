func moveZeroes(nums []int)  {
    nonZeroPos := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[nonZeroPos], nums[i] = nums[i], nums[nonZeroPos]
            nonZeroPos++
        }
    }
}