package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchUserEvents(username string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	fmt.Println("url: ", url)
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
	}

	var events []map[string]interface{}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&events); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return events, nil
}
