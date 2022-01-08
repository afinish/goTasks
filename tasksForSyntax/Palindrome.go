import "strings"
func isPalindrome(s string) bool {
    lowerCaseStr := strings.ToLower(s)
    
    var alphabetInt []uint8
    for i := 0; i < len(lowerCaseStr); i++ {
        if lowerCaseStr[i] >= 97 && lowerCaseStr[i] <= 122 {
            alphabetInt = append(alphabetInt, lowerCaseStr[i])
        }
    }
    
    for i, j := 0, 0; i < len(alphabetInt); i, j = i+1, j+1 {
        if alphabetInt[i] != alphabetInt[len(alphabetInt) - j - 1] {
            return false
        }
    }
    return true
}
