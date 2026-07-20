// Package variablelengthquantity implement simple library to perform
// variable length quantity encoding and decoding for uint32
package variablelengthquantity

import "errors"

// EncodeVarint returns encoded value
func EncodeVarint(input []uint32) []byte {
	output := []byte{}
	for _, n := range input {
		e := []byte{0x7F & byte(n)}
		for n = n >> 7; n > 0; n = n >> 7 {
			e = append([]byte{byte(0x7F&n) | 0x80}, e...)
		}
		output = append(output, e...)
	}
	return output
}

// DecodeVarint returns decoded value
func DecodeVarint(input []byte) ([]uint32, error) {
	if input[len(input)-1] > 0x7F {
		return []uint32{}, errors.New("incomplete sequence")
	}

	output := make([]uint32, 0)
	n := uint32(0)
	for _, b := range input {
		n = n<<7 + uint32(0x7F&b)
		if 0x80&b == 0 {
			output = append(output, n)
			n = 0
		}
	}
	return output, nil
}
