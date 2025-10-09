func sortArrayByParity(nums []int) []int {
    l := 0
    r := len(nums) - 1

    for l <= r {
        if nums[l] % 2 != 0 && nums[r] % 2 == 0 {
            nums[l], nums[r] = nums[r], nums[l]
        }

        if nums[l] % 2 == 0 {
            l++
        }

        if nums[r] % 2 != 0 {
            r--
        }
    }
    return nums
}