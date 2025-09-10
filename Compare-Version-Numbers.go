import "strconv"

func compareVersion(version1 string, version2 string) int {
    v1 := strings.Split(version1, ".")
    v2 := strings.Split(version2, ".")

    for i := 0; i < len(v1) || i < len(v2); i++ {
        var n1, n2 int
        if i < len(v1) {
            n1, _ = strconv.Atoi(v1[i])
        } 
        if i < len(v2) {
            n2, _ = strconv.Atoi(v2[i])
        }

        if n1 < n2 {
            return -1
        } 
        if n1 > n2 {
            return 1
        }
    }
    return 0
 }