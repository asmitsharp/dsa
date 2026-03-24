1func floodFill(image [][]int, sr int, sc int, color int) [][]int {
2
3    prevColor := image[sr][sc]
4    if color == prevColor {
5        return image
6    }
7    dfs(image, sr, sc, color, prevColor)
8
9    return image
10}
11
12func dfs(image [][]int, r, c , color, prevColor int) {
13    if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != prevColor {
14        return
15    }
16
17    image[r][c] = color
18
19    dfs(image, r+1, c, color, prevColor)
20    dfs(image, r-1, c, color, prevColor)
21    dfs(image, r, c+1, color, prevColor)
22    dfs(image, r, c-1, color, prevColor)
23} 
24