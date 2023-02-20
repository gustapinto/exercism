package series

func All(n int, s string) []string {
	var series []string

	for i, j := 0, n; j <= len(s); i, j = i+1, j+1  {
        series = append(series, s[i:j])
    }

    return series
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
