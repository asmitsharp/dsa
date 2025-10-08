func nextPermutation(nums []int)  {
    n := len(nums)
    k := -1
    l := -1

    for i := n - 2; i >= 0; i-- {
         if nums[i] < nums[i+1] {
            k = i
            break
        }
    }

    if k == -1 {
        reverse(nums, 0, n-1)
        return
    }

    for i := n - 1; i > k; i-- {
        if nums[i] > nums[k] {
            l = i
            break
        }
    }

    nums[k], nums[l] = nums[l], nums[k]

    reverse(nums, k+1, n-1)
}

func reverse(nums []int, start, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}