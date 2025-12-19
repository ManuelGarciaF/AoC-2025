package commons

import "strings"

type BitArray16 uint16

func (b BitArray16) Get(i int) bool {
	return (b & (1 << i)) != 0
}

func (b BitArray16) Set(i int) BitArray16 {
	return b | (1 << i)
}

func (b BitArray16) Clear(i int) BitArray16 {
	return b &^ (1 << i)
}

func (b BitArray16) Toggle(i int) BitArray16 {
	return b ^ (1 << i)
}

func (b BitArray16) String() string {
	var s strings.Builder
	s.WriteString("LSB:")
	for i := range 16 {
		if b&(1<<i) != 0 {
			s.WriteByte('1')
		} else {
			s.WriteByte('0')
		}
	}
	return s.String()
}
