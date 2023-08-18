// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
    romanMap := make(map[string]int)
    romanMap = map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }

    total := 0
    for i := 0; i < len(s); i++ {
        letterStr := string(s[i])
        value := romanMap[letterStr]
        if i == len(s) - 1 {
            total += value
            break
        }
        
        if romanMap[letterStr] < romanMap[string(s[i + 1])] {
            total -= value
        } else {
            total += value
        }
    }
    return total
}
