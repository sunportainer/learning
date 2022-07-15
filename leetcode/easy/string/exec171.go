package string

func titleToNumber(columnTitle string) int {
	ans := 0
	size := len(columnTitle)
	for i := 0; i < size; i++ {
		ans = ans*26 + int(columnTitle[i]-'A'+1)

	}
	return ans
}

//AB  = 26 + 2
