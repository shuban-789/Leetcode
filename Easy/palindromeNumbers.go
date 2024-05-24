import "strconv"
import "strings"

func reverse(s string) string {
    var output []string
    for j := len(s)-1; j >= 0; j-- {
        output = append(output, string(s[j]))
    }
    outstring := strings.Join(output, "")
    return outstring
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else if x < 10 {
        return true
    } else {
        xstring := strconv.Itoa(x)
        var xstringrev string
        xstringrev = reverse(xstring)
        if strings.Compare(xstring, xstringrev) == 0 {
            return true
        } else {
            return false
        }
    }
}
