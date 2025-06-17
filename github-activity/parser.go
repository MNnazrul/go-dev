package main

import "fmt"

func FormatEvent(event GitHubEvent) string {
	timeStr := TimeAgo(event.CreatedAt)
	base := fmt.Sprintf("[%s] ", timeStr)
	var msg string

	switch event.Type {
	case "PushEvent":
		commits := 0
		if c, ok := event.Payload["commits"].([]interface{}); ok {
			commits = len(c)
		}
		msg = fmt.Sprintf("Pushed %d commits to %s", commits, event.Repo.Name)
		return base + Colorize(msg, ColorGreen)
	case "WatchEvent":
		msg = fmt.Sprintf("Starred %s", event.Repo.Name)
		return base + Colorize(msg, ColorYellow)
	case "CreateEvent":
		refType := event.Payload["ref_type"]
		msg = fmt.Sprintf("Created a new %v in %s", refType, event.Repo.Name)
		return base + Colorize(msg, ColorCyan)
	case "ForkEvent":
		if forkee, ok := event.Payload["forkee"].(map[string]interface{}); ok {
			msg = fmt.Sprintf("Forked %s to %s", event.Repo.Name, forkee["full_name"])
		} else {
			msg = fmt.Sprintf("Forked %s", event.Repo.Name)
		}
		return base + Colorize(msg, ColorBlue)
	case "IssuesEvent":
		action := event.Payload["action"]
		msg = fmt.Sprintf("%v an issue in %s", action, event.Repo.Name)
		return base + Colorize(msg, ColorRed)
	case "IssueCommentEvent":
		msg = fmt.Sprintf("Commented on an issue in %s", event.Repo.Name)
		return base + Colorize(msg, ColorRed)
	case "PullRequestEvent":
		action := event.Payload["action"]
		msg = fmt.Sprintf("%v a pull request in %s", action, event.Repo.Name)
		return base + Colorize(msg, ColorBlue)
	case "PullRequestReviewCommentEvent":
		msg = fmt.Sprintf("Reviewed a pull request in %s", event.Repo.Name)
		return base + Colorize(msg, ColorBlue)
	case "DeleteEvent":
		refType := event.Payload["ref_type"]
		msg = fmt.Sprintf("Deleted a %v in %s", refType, event.Repo.Name)
		return base + Colorize(msg, ColorRed)
	default:
		msg = fmt.Sprintf("Unsupported event type: %s", event.Type)
		return base + msg
	}
}
