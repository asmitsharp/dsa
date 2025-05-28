import \sort\
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)

    n := len(nums)
    closestSum := nums[0] + nums[1] + nums[2]

    for i:= 0; i < n - 2; i++ {
        left, right := i + 1, n - 1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if abs(sum-target) < abs(closestSum - target) {
                closestSum = sum
            }

            if sum < target {
                left++
            } else if sum > target    {
                right--
            } else {
                return sum
            }
        }
    }
    return closestSum
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}