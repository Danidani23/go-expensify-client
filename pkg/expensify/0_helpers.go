package expensify

import (
	"fmt"
	"strings"
)

// splitFilenames takes a comma-separated string and returns a slice of filenames,
// or an error if any of the filenames are empty.
func splitFilenames(commaSeparatedFilenames string) ([]string, error) {
	filenames := strings.Split(commaSeparatedFilenames, ",")

	for i, filename := range filenames {
		filenames[i] = strings.TrimSpace(filename)
		if filenames[i] == "" {
			return nil, fmt.Errorf("invalid filename: filenames cannot be empty")
		}
	}

	return filenames, nil
}
