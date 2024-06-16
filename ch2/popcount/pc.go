package popcount

import (
	"crypto/sha256"
)

var p [256]byte

func init() {
	for i := range p {
		p[i] = p[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	// var res int
	// for i := 0; i < 63; i++ {
	// 	res += int(byte(x >> i & 1))
	// }
	// return res
	var res int
	for i := 0; i < 8; i++ {
		res += int(p[byte(x>>(i*8))])
	}
	return res
	// return int(
	// 	p[byte(x>>0*8)] +
	// 		p[byte(x>>1*8)] +
	// 		p[byte(x>>2*8)] +
	// 		p[byte(x>>3*8)] +
	// 		p[byte(x>>4*8)] +
	// 		p[byte(x>>5*8)] +
	// 		p[byte(x>>6*8)] +
	// 		p[byte(x>>7*8)])
}

func BitsDiff(b1, b2 *[sha256.Size]byte) int {
	var sum int
	for i := 0; i < sha256.Size; i++ {
		sum += int(p[b1[i]^b2[i]])
	}
	return sum
}
