package main

import (
	"fmt"
	"strings"
	"time"
)

func splitAndTrim(s string) []string {
	var out []string

	for _, part := range strings.Split(s, ",") {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func TimeAgo(isoTime string) string {
	t, err := time.Parse(time.RFC3339, isoTime)
	if err != nil {
		return ""
	}
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return fmt.Sprintf("%02d seconds ago", int(duration.Seconds()))
	case duration < time.Hour:
		return fmt.Sprintf("%02d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%02d hours ago", int(duration.Hours()))
	default:
		return fmt.Sprintf("%02d days ago", int(duration.Hours()/24))
	}
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
)

func Colorize(s, color string) string {
	return color + s + ColorReset
}

func isUnsupportedEvent(s string) bool {
	return len(s) >= 21 && s[21:41] == "Unsupported event type"
}
