package main

import "fmt"

func FormatEvent(event GitHubEvent) string {
	switch event.Type {
	case "PushEvent":
		commits := 0
		if c, ok := event.Payload["commits"].([]interface{}); ok {
			commits = len(c)
		}
		return fmt.Sprintf("Pushed %d commits to %s", commits, event.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("Starred %s", event.Repo.Name)
	case "CreateEvent":
		refType := event.Payload["ref_type"]
		return fmt.Sprintf("Created a new %v in %s", refType, event.Repo.Name)
	case "ForkEvent":
		if forkee, ok := event.Payload["forkee"].(map[string]interface{}); ok {
			return fmt.Sprintf("Forked %s to %s", event.Repo.Name, forkee["full_name"])
		}
		return fmt.Sprintf("Forked %s", event.Repo.Name)
	case "IssuesEvent":
		action := event.Payload["action"]
		return fmt.Sprintf("%v an issue in %s", action, event.Repo.Name)
	case "IssueCommentEvent":
		return fmt.Sprintf("Commented on an issue in %s", event.Repo.Name)
	case "PullRequestEvent":
		action := event.Payload["action"]
		return fmt.Sprintf("%v a pull request in %s", action, event.Repo.Name)
	case "PullRequestReviewCommentEvent":
		return fmt.Sprintf("Reviewed a pull request in %s", event.Repo.Name)
	case "DeleteEvent":
		refType := event.Payload["ref_type"]
		return fmt.Sprintf("Deleted a %v in %s", refType, event.Repo.Name)
	default:
		return fmt.Sprintf("Unsupported event type: %s", event.Type)
	}
}
