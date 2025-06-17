# GitHub Activity CLI

A command-line tool to fetch and display a user's most recent GitHub activity in a readable, colorful format.  
Inspired by the [GitHub User Activity project idea](https://roadmap.sh/projects/github-user-activity) from roadmap.sh.

---

## 🚀 Features

- Fetches recent public activity for any GitHub user.
- Friendly, natural language output with time formatting (e.g., `[2 hours ago]`).
- Supports filtering by event type and limiting the number of events.
- Caches API responses for 10 minutes to avoid rate limits and speed up repeated queries.
- Styled output with ANSI colors for different event types.
- No external libraries for HTTP or JSON—pure Go!

---

## 🛠️ Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/MNnazrul/github-activity.git
   cd github-activity
   ```

2. **Build the CLI:**
   ```sh
   go build -o github-activity
   ```

---

## ⚡ Usage

```sh
github-activity <github-username> [flags]
```

### **Flags**

| Flag               | Description                                       | Example                         |
| ------------------ | ------------------------------------------------- | ------------------------------- |
| `--filter` or `-f` | Filter by event types (comma-separated)           | `--filter=PushEvent,WatchEvent` |
| `--limit` or `-l`  | Number of events to display (max 100, default 10) | `--limit=5`                     |
| `--refresh`        | Skip cache and force API request                  | `--refresh`                     |
| `--raw`            | Show raw JSON output                              | `--raw`                         |

---

## 📝 **Examples**

Fetch the latest 10 events for a user:

```sh
github-activity MNnazrul
```

Show only push and star events, limit to 5:

```sh
github-activity MNnazrul --filter=PushEvent,WatchEvent --limit=5
```

Force a fresh API request (ignore cache):

```sh
github-activity MNnazrul --refresh
```

Show raw JSON output:

```sh
github-activity MNnazrul --raw
```

---

## 🎨 **Sample Output**

```
[2 hours ago] Pushed 3 commits to MNnazrul/go-backend
[5 hours ago] Starred kamranahmedse/developer-roadmap
[1 day ago] Created a new branch in MNnazrul/todo-cli
```

_Output is colorized by event type if your terminal supports ANSI colors._

---

## 🗂️ **Supported Event Types**

- PushEvent
- WatchEvent
- CreateEvent
- ForkEvent
- IssuesEvent
- IssueCommentEvent
- PullRequestEvent
- PullRequestReviewCommentEvent
- DeleteEvent

_Unknown event types are skipped with a notice._

---

## 💾 **Caching**

- API responses are cached in `.github-activity-cache.json` in your home directory.
- Cache is valid for 10 minutes.
- Use `--refresh` to ignore cache and fetch fresh data.

---

## ⚠️ **GitHub API Rate Limits**

- Unauthenticated requests are limited by GitHub (usually 60 requests/hour).
- If you hit the rate limit, the CLI will show a friendly message.
- Caching helps avoid unnecessary API calls.

---

## 📚 **Learning Outcomes**

By building and using this tool, you’ll learn:

- How to interact with REST APIs in Go (without external libraries)
- How to parse and format JSON data
- How to build modular CLI tools in Go
- How to implement caching for CLI tools
- How to handle errors and edge cases gracefully

---

## 📎 **Project Inspiration**

This project idea is from [roadmap.sh: GitHub User Activity](https://roadmap.sh/projects/github-user-activity).

---

## 📝 **License**

MIT License
