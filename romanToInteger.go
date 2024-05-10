func romanToInt(s string) int {
    numeralValue := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
    sum := 0
    for i := 0; i < len(s); i++ {
        if i+1 < len(s) {
            if numeralValue[string(s[i+1])] > numeralValue[string(s[i])] {
                sum += (numeralValue[string(s[i+1])] - numeralValue[string(s[i])])
                i += 1
            } else {
                sum += numeralValue[string(s[i])]
            }
        } else {
            sum += numeralValue[string(s[i])]
        }
    }
    return sum
}
