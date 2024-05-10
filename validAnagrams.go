import (
"slices"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    } else {
        var t_array, s_array []string
        for i := 0; i < len(s); i++ {
            s_array = append(s_array, string(s[i]))
            t_array = append(t_array, string(t[i]))
        }
        slices.Sort(s_array)
        slices.Sort(t_array)
        for n := 0; n < len(s_array); n++ {
            if s_array[n] != t_array[n] {
                return false
            }
        }
        return true
    }
}
