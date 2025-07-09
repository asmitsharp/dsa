

func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
        return ""
    }

    tFreq := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        tFreq[t[i]]++
    }

    required := len(tFreq) 
    formed := 0           
    windowFreq := make(map[byte]int)
    start, minLen := 0, len(s)+1
    resultStart, resultEnd := 0, 0

    for right := 0; right < len(s); right++ {
    
        char := s[right]
        windowFreq[char]++

        
        if count, needed := tFreq[char]; needed && windowFreq[char] == count {
            formed++
        }

    
        for left := start; formed == required && left <= right; left++ {

            if right-left+1 < minLen {
                minLen = right - left + 1
                resultStart, resultEnd = left, right+1
            }

           
            char = s[left]
            windowFreq[char]--
            if count, needed := tFreq[char]; needed  && windowFreq[char] < count {
                formed--
            }
            start++
        }
    }

 
    if minLen > len(s) {
        return ""
    }
    return s[resultStart:resultEnd]
}