package mymath

//Even 判断正整数奇偶性,是否偶数
func Even(number uint64) bool {
	return number%2 == 0
}

//Odd 判断正整数奇偶性，是否奇数
func Odd(number uint64) bool {
	// Odd should return not even.
	// ... We cannot check for 1 remainder.
	// ... That fails for negative numbers.
	return !Even(number)
}

//Floor 地板函数
func Floor(x float64) int {
	if x < 0 {
		return int(x) - 1
	}
	return int(x)
}

//IsPrime ...
func IsPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
