// Package accumulate implements a simple library that
// performs an operation on a collection of strings
package accumulate

// Accumulate performs given operation on each memeber of the input collection
func Accumulate(collection []string, operation func(string) string) []string {
	for i, element := range collection {
		collection[i] = operation(element)
	}
	return collection
}
