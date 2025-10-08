func maxArea(height []int) int {
    left := 0
    right := len(height) - 1

    max_area := 0

    for left < right {
        width := right - left
        curr_area := 0
        if height[left] < height[right] {
            curr_area = width * height[left]
            left++
        } else {
            curr_area = width * height[right]
            right--
        }

        if curr_area > max_area {
            max_area = curr_area
        }
    }

    return max_area
}