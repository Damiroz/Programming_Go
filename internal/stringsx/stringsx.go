package stringsx

// Clip возвращает строку s, обрезанную до длины max
func Clip(s string, max int) string {
	if max < 0 {
		max = 0
	}
	if len(s) <= max {
		return s
	}
	return s[:max]
}