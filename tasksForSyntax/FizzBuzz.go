import "strconv"
func fizzBuzz(n int) []string {
    res := make([]string, n)
    for i := 1; i <= n; i++ {
        res[i - 1] = FB(i)
    }
    return res
}

func FB(n int) string {
        if (n % 3 == 0) && (n % 5 == 0) {
        return "FizzBuzz"
    }
    if n % 3 == 0 {
        return "Fizz"
    }
    if n % 5 == 0 {
        return "Buzz"
    }
    return strconv.Itoa(n)
}
