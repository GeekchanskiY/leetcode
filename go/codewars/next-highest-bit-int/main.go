package main

func intToBits(n int) []int {
	var bits []int
	for n > 0 {
		bit := n & 1
		bits = append(bits, bit)
		n = n >> 1
	}
	return bits
}

func bitsToInt(bits []int) int {
	n := 0
	for i := len(bits) - 1; i >= 0; i-- {
		n = (n << 1) | bits[i]
	}
	return n
}

func NextHigher(n int) int {
	bits := intToBits(n)
	bits = append(bits, 0)
	for i := len(bits) - 1; i > 0; i-- {
		println(bits[i])
		if bits[i] == 0 && bits[i-1] == 1 {
			bits[i] = 1
			bits[i-1] = 0
			break
		}
	}
	return bitsToInt(bits)
}

func main() {
	println(NextHigher(127))
}
