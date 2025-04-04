package types

type ByteView []byte

func (b ByteView) Len() int {
	return len(b)
}

func (v ByteView) ByteSlice() []byte {
	return CloneBytes(v)
}

func (b ByteView) String() string {
	return string(b)
}

func CloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
