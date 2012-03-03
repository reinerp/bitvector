package bitvector

const (
	width = 8
	shift = 3
	mask = 0x7
    )

type Bitvector struct {
	data []uint8
	length int
}

// Construct the vector with length len, capacity  and fill with the value b
func New(length, capacity int) *Bitvector {
	return &Bitvector{ make([]uint8, size(length), size(capacity)), length }
}

// Get the i'th element
func (b Bitvector) At(i int) bool {
	b.checkLength(i)
	return (b.data[i >> shift] & (1<< uint8(i & mask))) != 0
}

// Set the i'th element
func (b Bitvector) Set(i int, val bool) {
	b.checkLength(i)
	if val {
		b.data[i >> shift] |= (1 << uint8(i & mask))
	} else {
		b.data[i >> shift] &= ^(1 << uint8(i & mask))
	}
}

func size(length int) int {
	res := length/width
	if 0 != length & mask {
		res++
	}
	return res
}

func (b Bitvector) checkLength(i int) {
/*	if i < 0 || i >= b.length {
		panic("bitvector index out of range")
	}*/
}
