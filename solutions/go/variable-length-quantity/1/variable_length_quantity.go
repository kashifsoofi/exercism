// Package variablelengthquantity implement simple library to perform
// variable length quantity encoding and decoding for uint32
package variablelengthquantity

import "errors"

// EncodeVarint returns encoded value
func EncodeVarint(input []uint32) []byte {
	output := []byte{}
	for _, n := range input {
		output = append(output, encode(n)...)
	}
	return output
}

// DecodeVarint returns decoded value
func DecodeVarint(input []byte) ([]uint32, error) {
	if input[len(input)-1] > 127 {
		return []uint32{}, errors.New("incomplete sequence")
	}

	output := make([]uint32, 0)
	n := uint32(0)
	for _, b := range input {
		n = n << 7
		n += uint32(b & 0x7F)
		if b&0x80 == 0 {
			output = append(output, n)
			n = 0
		}
	}
	return output, nil
}

func encode(n uint32) []byte {
	e := []byte{0x7F & byte(n)}
	n = n >> 7
	for n > 0 {
		b := byte(0x7F&n) | 0x80
		e = append([]byte{b}, e...)
		n = n >> 7
	}
	return e
}
