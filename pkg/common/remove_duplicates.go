package common

func RemoveDuplicates(elements []string) []string {
	// Create a map to track seen elements
	seen := make(map[string]struct{})
	result := []string{}

	// Loop over the elements
	for _, element := range elements {
		// Check if the element is already seen
		if _, found := seen[element]; !found {
			// If not seen, add it to the result and mark it as seen
			result = append(result, element)
			seen[element] = struct{}{}
		}
	}

	return result
}
