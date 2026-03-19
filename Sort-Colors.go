1func sortColors(nums []int)  {
2    l := 0
3    curr := 0
4    r := len(nums) - 1
5
6    for curr <= r {
7        switch nums[curr] {
8            case 0: 
9             nums[l], nums[curr] = nums[curr] , nums[l]
10             l++
11             curr++
12            case 1:
13                curr++
14            case 2:
15                nums[r], nums[curr] = nums[curr], nums[r]
16                r--
17        }
18    }
19}