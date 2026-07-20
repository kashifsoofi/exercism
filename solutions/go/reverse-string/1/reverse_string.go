// Package reverse implements a simple library that reverse a string
package reverse

// Reverse returns reversed string, if successful.
func Reverse(input string) string {
	arr := []rune(input)

	length := len(arr)
	limit := length / 2
	ri := length - 1
	for i := 0; i < limit; i++ {
		a := arr[i]
		arr[i] = arr[ri]
		arr[ri] = a
		ri--
	}

	return string(arr)
}
