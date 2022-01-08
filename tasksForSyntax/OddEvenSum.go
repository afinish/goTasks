func oddEvenSum(arr []int) (int, int) {
	var odd, even int
	odd, even = 0, 0
	
	for _, key := range arr {
		if key % 2 == 0 {
			even = even + key
		} else {
			odd = odd + key
		}
	}
	return odd, even
}
