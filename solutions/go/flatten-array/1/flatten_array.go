// Package flatten implements simple routing to flatten nested list
package flatten

// Flatten flattens nested list of objects and returns
func Flatten(input interface{}) []interface{} {
	output := make([]interface{}, 0)
	if arr, ok := input.([]interface{}); ok {
		for _, i := range arr {
			output = append(output, Flatten(i)...)
		}
	} else if input != nil {
		output = append(output, input)
	}

	return output
}
