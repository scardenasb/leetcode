// https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		letter := string(strs[0][i])
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || string(strs[j][i]) != letter {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
