package medium

import "fmt"

func maxPoints(points [][]int) int {

	myabs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	gcd := func(a, b int) int {
		a = myabs(a)
		b = myabs(b)
		if a < b {
			tmp := a
			a = b
			b = tmp
		}
		for b != 0 {
			m := a % b
			a = b
			b = m
		}
		return a
	}
	size := len(points)
	if 0 == size {
		return 0
	}

	maxp := 0
	for i := 0; i < size; i++ {
		hashmap := make(map[string]int, size)
		res := 0
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			x := points[i][0] - points[j][0]
			y := points[i][1] - points[j][1]

			g := gcd(x, y)
			x = x / g
			y = y / g
			key := fmt.Sprintf("%d@%d", x, y)
			hashmap[key] = hashmap[key] + 1
			if res < hashmap[key] {
				res = hashmap[key]
			}
		}
		if maxp < res {
			maxp = res
		}
	}
	return maxp + 1

}
