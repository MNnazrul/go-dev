package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchUserEvents(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	fmt.Println("url: ", url)
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		// continue
	case http.StatusNotFound:
		return nil, fmt.Errorf("user '%s' not found(404)", username)
	case http.StatusForbidden:
		return nil, fmt.Errorf("API rate limit exceeded or access forbidden (403)")
	default:
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s\n%s", resp.Status, string(body))
	}

	var events []GitHubEvent

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&events); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return events, nil
}
