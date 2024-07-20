package common

func Contains[T comparable](list []T, element T) bool {
	for _, curr := range list {
		if curr == element {
			return true
		}
	}
	return false
}
