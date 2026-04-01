1func mySqrt(x int) int {
2    lo , hi := 0, x
3
4    for lo <= hi {
5        mid := lo + (hi - lo) / 2
6
7        sq := mid * mid
8        if sq == x {
9            return mid
10        } else if sq < x {
11            lo = mid + 1
12        } else {
13            hi = mid - 1
14        }
15    }
16    return hi
17}