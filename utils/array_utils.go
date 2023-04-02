package utils

func Filter[T any](slice []T, f func(T) bool) []T {
	var output []T
	for _, element := range slice {
		if f(element) {
			output = append(output, element) //the values that satisfy filter will be appended in the output
		}
	}
	return output
}
