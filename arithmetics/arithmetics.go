package arithmetics

// Pow implements binary power algo
func Pow(a, b int64) int64 {
	var res int64 = 1
	for b > 0 {
		if b&1 != 0 {
			res *= a
		}
		a *= a
		b >>= 1
	}
	return res
}
