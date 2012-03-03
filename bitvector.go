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

func (b Bitvector) checkLength(i int) {
	if i < 0 || i >= b.length {
		panic("bitvector index out of range")
	}
}

func (b Bitvector) Get(i int) bool {
	b.checkLength(i)
	return (b.data[i >> shift] & (1<< uint8(i & mask))) != 0
}

func (b Bitvector) Set(i int, val bool) {
	b.checkLength(i)
	if val {
		b.data[i >> shift] |= (1 << uint8(i & mask))
	} else {
		b.data[i >> shift] &= ^(1 << uint8(i & mask))
	}
}