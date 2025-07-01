func predictPartyVictory(senate string) string {
    n := len(senate)
    radiant := []int{}
    dire := []int{}

    for  i, c := range senate {
        if c == 'R' {
            radiant = append(radiant, i)
        } else {
            dire = append(dire, i)
        }
    }

    for len(radiant) > 0 && len(dire) > 0 {
        if radiant[0] < dire[0] {
            radiant = append(radiant, radiant[0] + n)
        } else {
            dire = append(dire, dire[0] + n)
        }

        radiant = radiant[1:]
        dire = dire[1:]
    }

    if len(radiant) > 0 {
        return "Radiant"
    } 

    return "Dire"
}