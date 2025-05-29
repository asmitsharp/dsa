func sortColors(nums []int)  {
   for i := 0; i < len(nums) - 1; i++ {
        minIndex := i
        for j := i + 1; j < len(nums); j++ {
            if nums[minIndex] > nums[j] {
                minIndex = j
            }
        }

        nums[i], nums[minIndex] = nums[minIndex], nums[i]
    } 
}

