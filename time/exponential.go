package time

func Exponential(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}
