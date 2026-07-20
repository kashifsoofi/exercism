// Package accumulate implements a simple library that
// performs an operation on a collection of strings
package accumulate

// Accumulate returns true if given word is an isogram, if successful
func Accumulate(collection []string, operation func(string) string) []string {
	for i, element := range collection {
		collection[i] = operation(element)
	}
	return collection
}
