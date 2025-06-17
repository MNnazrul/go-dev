package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const (
	cacheFileName = ".github-activity-cache.json"
	cacheDuration = 10 * time.Minute
)

type CacheEntry struct {
	Timestamp time.Time     `json:"timestamp"`
	Events    []GitHubEvent `json:"events"`
}

type Cache map[string]CacheEntry

func getCacheFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return cacheFileName
	}
	return filepath.Join(home, cacheFileName)
}

func LoadCache() (Cache, error) {
	path := getCacheFilePath()
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Cache{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var cache Cache
	if err := json.NewDecoder(file).Decode(&cache); err != nil {
		return nil, err
	}
	return cache, nil
}

func SaveCache(cache Cache) error {
	path := getCacheFilePath()
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	return encoder.Encode(cache)
}

func GetCachedEvents(cache Cache, username string) ([]GitHubEvent, bool) {
	entry, ok := cache[username]
	if !ok {
		return nil, false
	}
	if time.Since(entry.Timestamp) > cacheDuration {
		return nil, false
	}
	return entry.Events, true
}

func SetCachedEvents(cache Cache, username string, events []GitHubEvent) {
	cache[username] = CacheEntry{
		Timestamp: time.Now(),
		Events:    events,
	}
}
