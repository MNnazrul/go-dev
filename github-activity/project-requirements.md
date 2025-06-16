## ✅ Complete Project Requirements: GitHub Activity CLI

---

### 1. 🎯 **Project Objective**

Create a command-line tool `github-activity` that:

- Takes a GitHub username as an argument.
- Fetches and displays the user's most recent GitHub activity in a readable format.
- Includes helpful CLI features like filters, limits, caching, etc.
- Must be implemented in **Go (Golang)** without using external libraries for HTTP or JSON.

---

### 2. 🧱 **Core Features**

#### 2.1. Run from the command line

- Command format:

  ```bash
  github-activity <username>
  ```

#### 2.2. Fetch Data

- API endpoint:

  ```
  https://api.github.com/users/<username>/events
  ```

#### 2.3. Supported Event Types to Display

You must handle and display the following GitHub event types, with friendly output:

| Event Type                      | Output Example                                                 |
| ------------------------------- | -------------------------------------------------------------- |
| `PushEvent`                     | Pushed X commits to `<repo>`                                   |
| `WatchEvent`                    | Starred `<repo>`                                               |
| `CreateEvent`                   | Created a new `<ref_type>` (`branch`/`repository`) in `<repo>` |
| `ForkEvent`                     | Forked `<repo>` to `<forkee.full_name>`                        |
| `IssuesEvent`                   | Opened/Closed an issue in `<repo>`                             |
| `IssueCommentEvent`             | Commented on an issue in `<repo>`                              |
| `PullRequestEvent`              | Opened/Closed/Merged a pull request in `<repo>`                |
| `PullRequestReviewCommentEvent` | Reviewed a pull request in `<repo>`                            |
| `DeleteEvent`                   | Deleted a `<ref_type>` in `<repo>`                             |

**Note**: Skip unknown event types with a message like "Unsupported event type: <type>"

#### 2.4. Format and Output

- Display the most recent 10 events (by default), newest at the top.
- Format each activity in **natural language**, e.g.:

  ```
  [2 hours ago] Pushed 3 commits to kamranahmedse/developer-roadmap
  [5 hours ago] Starred kamranahmedse/developer-roadmap
  [1 day ago] Opened a new issue in kamranahmedse/developer-roadmap
  ```

---

### 3. 📥 **Optional Arguments (Flags)**

All optional arguments should be processed via standard CLI flag handling.

| Flag               | Description                                                | Example                         |
| ------------------ | ---------------------------------------------------------- | ------------------------------- |
| `--filter` or `-f` | Filter by one or more event types (comma-separated)        | `--filter=PushEvent,WatchEvent` |
| `--limit` or `-l`  | Set how many events to fetch and display (max 100)         | `--limit=20`                    |
| `--refresh`        | Skip cache and force API request                           | `--refresh`                     |
| `--raw`            | Show raw JSON of the API response (for debugging/learning) | `--raw`                         |

#### Valid event types for `--filter`:

```
PushEvent, WatchEvent, CreateEvent, ForkEvent, IssuesEvent, IssueCommentEvent,
PullRequestEvent, PullRequestReviewCommentEvent, DeleteEvent
```

---

### 4. 💥 **Error Handling**

You must handle these errors gracefully:

| Situation                          | Expected Behavior                                 |
| ---------------------------------- | ------------------------------------------------- |
| Invalid or empty username          | Print error and exit                              |
| No events available                | Inform the user                                   |
| GitHub user does not exist (404)   | Show a user-friendly message                      |
| API rate limit exceeded (403)      | Show message and optionally remaining limit/time  |
| Network error (timeout, DNS, etc.) | Show technical reason                             |
| JSON parsing error                 | Print technical error and suggest retry           |
| Unsupported event type             | Skip it with a notice if `--verbose` flag is used |

---

### 5. 💾 **Caching Feature**

- Store the last successful API response in a local file (`cache_<username>.json`).
- Use this cached file if:

  - The user doesn't include `--refresh`
  - The cache is not older than 10 minutes.

- Print: `Using cached data (updated 5 minutes ago)` if using cache.

---

### 6. 🖍️ **Bonus: Styled Output**

If terminal supports color (you can detect via `os.Stdout`):

- Use ANSI escape codes for styling:

  - `PushEvent` → Green
  - `IssuesEvent`, `IssueCommentEvent`, `PullRequestEvent` → Blue
  - `ForkEvent`, `CreateEvent` → Cyan
  - `WatchEvent` → Yellow

Example:

```
[2h ago] \033[32mPushed 3 commits\033[0m to kamranahmedse/developer-roadmap
```

---

### 7. 📦 **Directory Structure Recommendation**

```
github-activity/
├── main.go                  # Entry point
├── api.go                   # Handles API calls
├── parser.go                # JSON parsing and formatting
├── cache.go                 # Caching logic
├── flags.go                 # CLI argument handling
├── utils.go                 # Helpers: colors, error, time formatting
├── types.go                 # Structs for event types
├── go.mod / go.sum          # Go module files
├── README.md                # Usage & documentation
├── /test/                   # Unit tests
```

---

### 8. 🧪 **Testing Requirements**

- Unit tests for:

  - Fetching from the API
  - Parsing and formatting various event types
  - Handling edge cases (empty events, unknown event type)
  - CLI argument parsing logic

---

### 9. 📘 **README.md Should Include**

- Purpose of the tool
- How to build & run
- Usage examples with all flags
- Sample outputs
- Explanation of supported event types
- Notes on GitHub API rate limits
- How caching works
- Screenshots or terminal recordings (optional)

---

### 10. 🎓 **Learning Outcomes**

By completing this project, you'll learn:

- How to interact with REST APIs using raw Go (no external libs)
- How to parse nested JSON into Go structs
- How to format time differences (e.g., “2 hours ago”)
- How to build command-line interfaces in Go
- How to write modular, testable Go code
- How to cache and reuse network responses
- How to build a tool with real-world value

---
