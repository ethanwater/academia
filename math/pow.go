package math

func Pow(x float64, n int) float64 {
	pospow := func(x float64, n int) float64 {
		initial := x
		for i := 1; i < n; i++ {
			x = x * initial
		}
		return x
	}
	switch {
	case n > 0:
		return pospow(x, n)
	case n < 0:
		return float64(1) / pospow(x, -n)
	default:
		return float64(1)
	}
}

func FastPow(x float64, n int) float64 {
	pospow := func(x float64, n int) float64 {
		init := 1.0
		a := x
		for n > 0 {
			if n&1 != 0 {
				init *= a
			}
			a *= a
			n >>= 1
		}
		return init
	}
	switch {
	case n > 0:
		return pospow(x, n)
	case n < 0:
		return 1.0 / pospow(x, -n)
	default:
		return 1.0
	}
}
