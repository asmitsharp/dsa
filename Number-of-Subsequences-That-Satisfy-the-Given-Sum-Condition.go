func numSubseq(nums []int, target int) int {
    sort.Ints(nums)

    const mod = 1_000_000_007
    pow2 := make([]int64, len(nums)+1)
    pow2[0] = 1
    for i := 1; i <= len(nums); i++ {
        pow2[i] = (pow2[i-1] * 2) % mod
    }

    l, r := 0, len(nums) - 1
    var count int64
    
    for l <= r {
        sum := nums[l] + nums[r]
        if sum <= target {
            count = (count + pow2[r-l]) % mod
            l++
        } else if sum > target {
            r--
        }
    }
    return int(count)
}