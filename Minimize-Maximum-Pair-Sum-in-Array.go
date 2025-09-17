func minPairSum(nums []int) int {
    sort.Ints(nums)
    l, r := 0, len(nums) - 1
    maxSum := 0
    for l < r {
        sum := nums[l] + nums[r] 
        if sum > maxSum {
            maxSum = sum
        }
        l++
        r--
    }
    return maxSum
}