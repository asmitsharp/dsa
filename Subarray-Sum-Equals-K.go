func subarraySum(nums []int, k int) int {
    count := 0
    prefixSum := 0

    prefixCount := map[int]int{0: 1}

    for _, num := range nums {
        prefixSum += num

        if val, exists := prefixCount[prefixSum - k]; exists {
            count += val
        }

        prefixCount[prefixSum]++
    }

    return count
}

