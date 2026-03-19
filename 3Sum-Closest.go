1func threeSumClosest(nums []int, target int) int {
2    sort.Ints(nums)
3    closestSum := nums[0] + nums[1] + nums[2]
4
5    for i := 0; i < len(nums) - 2; i++ {
6        l := i+1
7        r := len(nums) - 1
8
9        for l < r {
10            sum := nums[i] + nums[l] + nums[r]
11
12            if sum == target {
13                return sum
14            }
15
16            if abs(sum-target) < abs(closestSum-target) {
17                closestSum = sum
18            }
19
20            if sum < target {
21                l++
22            } else {
23                r--
24            }
25
26        }
27    }
28
29    return closestSum
30}
31
32func abs(num int) int {
33    if num < 0 {
34        return -num
35    }
36    return num
37}