import "sort"

func findContentChildren(g []int, s []int) int {
    sort.Ints(s)
    sort.Ints(g)

    count := 0
    for j := 0; j < len(g); j++ {
        for i := 0; i < len(s); i++ {
            if s[i] >= g[j] {
                count += 1
                s = append(s[:i], s[i+1:]...)
                break
            }
        }
    }
    return count
}
