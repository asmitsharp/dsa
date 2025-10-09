func flipAndInvertImage(image [][]int) [][]int {
    
    for _, row := range image {
        l := 0
        r := len(row) - 1

        for l <= r {
            old := row[l]
            row[l] = 1 - row[r]
            row[r] = 1 - old
            l++
            r--
        }
    }
    return image
}