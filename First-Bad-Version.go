1/** 
2 * Forward declaration of isBadVersion API.
3 * @param   version   your guess about first bad version
4 * @return 	 	      true if current version is bad 
5 *			          false if current version is good
6 * func isBadVersion(version int) bool;
7 */
8
9func firstBadVersion(n int) int {
10    lo , hi := 1, n
11
12    for lo < hi {
13        mid := lo + (hi - lo) / 2
14
15        if isBadVersion(mid) {
16            hi = mid
17        } else {
18            lo = mid + 1
19        }
20    }
21    return lo
22}