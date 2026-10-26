// Package solution contains the Go solution for LeetCode 125.
package solution

func isPalindrome(s string) bool {

	i := 0
	j := len(s) - 1

	for i < j {
		if !isASCIIAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isASCIIAlphanumeric(s[j]) {
			j--
			continue
		}
		if toLowerASCII(s[i]) != toLowerASCII(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLowerASCII(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isASCIIAlphanumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
