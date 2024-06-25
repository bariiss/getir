package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// Define the URL
	var url = "http://ip-api.com/json"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Unmarshal the response body into a map
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Beautify the JSON output
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert pretty JSON to string
	jsonString := string(prettyJSON)

	// Color and bold the JSON keys
	coloredJSON := colorizeJSON(jsonString)

	// Print the beautified and colored JSON
	fmt.Println(coloredJSON)
}

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
