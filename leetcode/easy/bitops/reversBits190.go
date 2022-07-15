package bitops

func reverseBits(num uint32) uint32 {
	var ret uint32
	ret = 0
	for i := 1; i <= 32; i++ { //loop 32 times
		tmp := num & 1
		ret = ret << 1
		ret = ret | tmp
		num = num >> 1
	}
	return ret

}
