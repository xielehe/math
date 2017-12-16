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
func Floor(x float32) int {
	if x < 0 {
		return int(x) - 1
	}
	return int(x)
}
