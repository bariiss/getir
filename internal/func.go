package internal

import (
	"fmt"
	"strings"
)

func colorizeJSON(jsonString string) string {
	lines := strings.Split(jsonString, "\n")
	for i, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			key := parts[0]
			value := strings.Join(parts[1:], ":")
			// Apply color and bold to the key
			coloredKey := fmt.Sprintf("\033[1;34m%s\033[0m", key)
			lines[i] = coloredKey + ":" + value
		}
	}
	return strings.Join(lines, "\n")
}
