package time

func Logarithmic(n int) int {
	m := 0
	i := 1
	for i < n {
		m += 1
		i = i * 2
	}
	return m
}
