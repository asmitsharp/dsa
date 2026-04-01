1/** 
2 * Forward declaration of guess API.
3 * @param  num   your guess
4 * @return 	     -1 if num is higher than the picked number
5 *			      1 if num is lower than the picked number
6 *               otherwise return 0
7 * func guess(num int) int;
8 */
9
10func guessNumber(n int) int {
11    lo , hi := 0, n
12
13    for true {
14        mid := lo + (hi - lo) / 2
15
16        guess := guess(mid)
17
18        if guess == 1 {
19            lo = mid + 1
20        } else if guess == -1 {
21            hi = mid - 1
22        } else if guess == 0 {
23            return mid
24        }
25    }
26    return -1
27}