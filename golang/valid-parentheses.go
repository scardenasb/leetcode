// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	charsMap := make(map[rune]rune)
	charsMap[')'] = '('
	charsMap[']'] = '['
	charsMap['}'] = '{'

	newStr := []rune(s)
	if len(newStr) == 0 {
		return false
	}

	for i := 0; i < len(newStr); i++ {
		if newStr[i] == ')' || newStr[i] == ']' || newStr[i] == '}' {
			if i == 0 || newStr[i-1] != charsMap[newStr[i]] {
				return false
			} else {
				newStr = append(newStr[:i-1], newStr[i+1:]...)
				i -= 2
			}
		}
	}
	return len(newStr) == 0
}
