1func findMinArrowShots(points [][]int) int {
2    if len(points) == 0 {
3        return 0
4    }
5
6    // sort by end
7    sort.Slice(points, func(i, j int) bool {
8        return points[i][1] < points[j][1]
9    })
10
11    arrows := 1
12    currentEnd := points[0][1]
13
14    for i := 1; i < len(points); i++ {
15        if points[i][0] > currentEnd {
16            // not burst by this arrow
17            arrows++
18            currentEnd = points[i][1]
19        }
20    }
21
22    return arrows
23}