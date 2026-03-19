1func trap(height []int) int {
2    left := 0
3    right := len(height) - 1
4    maxLeft, maxRight := 0,0
5    water := 0
6
7    for left < right {
8        if height[left] <= height[right] {
9            if height[left] >= maxLeft {
10                maxLeft = height[left]
11            } else {
12                water += maxLeft - height[left] 
13            }
14            left += 1
15        } else {
16            if height[right] >= maxRight {
17                maxRight = height[right]
18            } else {
19                water +=  maxRight - height[right]
20            }
21            right -= 1
22        }
23    }
24
25    return water
26}