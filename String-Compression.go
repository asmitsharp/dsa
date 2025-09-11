func compress(chars []byte) int {
    write := 0
    read := 0
    n := len(chars)

    for read < n {
        count := 0
        cur := chars[read] 

        for read < n && chars[read] == cur {
            read++
            count++
        }

        chars[write] = cur
        write++
        
        if count > 1 {
            for _, d := range strconv.Itoa(count) {
                chars[write] = byte(d)
                write++
            }
        }
    }

    return write
}