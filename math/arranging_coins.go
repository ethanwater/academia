package math

func arrange(n int) int {
	if n <= 1 {
        return n
    }
    slots := 0
    for i := 1; i <= n; i++ {
        slots += i
        if slots > n {
            return i - 1
        }
    }
    return 0
}
