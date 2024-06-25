package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bariiss/getir/internal"
)

func main() {
	// Define the URL
	url := "http://ip-api.com/json"

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
	coloredJSON := internal.ColorizeJSON(jsonString)

	// Print the beautified and colored JSON
	fmt.Println(coloredJSON)
}
