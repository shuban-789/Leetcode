import "strings"

func lengthOfLastWord(s string) int {
    var spaceTolerance int = 0
    var length int = 0
    for i := len(s) - 1; i > -1; i-- {
        if strings.Compare(string(s[i]), " ") != 0 {
            length++
            spaceTolerance++
        } else if (spaceTolerance == 0) {
            continue
        } else {
            break
        }
    }
    return length
}

// Note: Excellent runtime
