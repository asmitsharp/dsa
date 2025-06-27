/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */

func deserialize(s string) *NestedInteger {
    if len(s) == 0 {
        return &NestedInteger{}
    }

    // Case: single integer
    if s[0] != '[' {
        num, _ := strconv.Atoi(s)
        ni := NestedInteger{}
        ni.SetInteger(num)
        return &ni
    }

    stack := []*NestedInteger{}
    num := 0
    negative := false
    hasNum := false

    for i := 0; i < len(s); i++ {
        ch := s[i]

        if ch == '-' {
            negative = true
        } else if ch >= '0' && ch <= '9' {
            num = num*10 + int(ch-'0')
            hasNum = true
        } else if ch == '[' {
            ni := NestedInteger{}
            stack = append(stack, &ni)
        } else if ch == ',' || ch == ']' {
            if hasNum {
                if negative {
                    num = -num
                }
                ni := NestedInteger{}
                ni.SetInteger(num)
                stack[len(stack)-1].Add(ni)
            }
            // Reset for next number
            num = 0
            negative = false
            hasNum = false

            if ch == ']' && len(stack) > 1 {
                ni := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                stack[len(stack)-1].Add(*ni)
            }
        }
    }

    return stack[0]
}
