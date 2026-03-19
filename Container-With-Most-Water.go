1func maxArea(height []int) int {
2    left := 0
3    right := len(height) - 1
4    maxWater := 0
5
6    for left < right {
7        water := 0
8        if height[left] > height[right] {
9            water = height[right] * (right - left)
10            right--
11        } else {
12            water = height[left] * (right - left)
13            left++
14        }
15
16        if water > maxWater {
17            maxWater = water
18        }
19    }
20
21    return maxWater
22}