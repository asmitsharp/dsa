1func merge(nums1 []int, m int, nums2 []int, n int)  {
2    p1, p2, p := m-1, n-1, m+n-1
3
4    for p2 >= 0 {
5        if p1 >= 0 && nums1[p1] > nums2[p2] {
6            nums1[p] = nums1[p1]
7            p1--
8        } else {
9            nums1[p] = nums2[p2]
10            p2--
11        }
12        p--
13    }
14}