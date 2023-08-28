package file

import (
	"bufio"
	"io"
	"strings"
)

// ReadYamlComments reads the continous comments from a yaml file, if a non-comment line is found, it stops reading.
func ReadYamlComments(r io.Reader) (string, error) {
	// create a new scaner
	scanner := bufio.NewScanner(r)

	content := ""

	// scan the file
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "# ") {
			// replace the first 2 characters with empty string
			line = strings.Replace(line, "# ", "", 1)
			content += line + "\n"
		} else {
			break
		}
	}

	return content, scanner.Err()
}
