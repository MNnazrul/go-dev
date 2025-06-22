# ✅ Project Name: **Caching Proxy CLI Tool**

---

## 🔧 Overview

Build a **Command-Line Interface (CLI) tool** that launches a **caching reverse proxy server**. The proxy should forward incoming HTTP requests to a specified **origin server**, cache the responses, and serve cached responses for repeated requests.

---

## 📌 Core Requirements

### 1. **Start the Proxy Server via CLI**

#### ✅ Feature

The user can start the proxy server by running:

```bash
caching-proxy --port <number> --origin <url>
```

#### 📍 What to Implement

- Parse command-line arguments:

  - `--port`: required, number. The port where your proxy server will listen.
  - `--origin`: required, URL. The base origin server where requests are forwarded.

- Start an HTTP server on the specified port.

---

### 2. **Request Forwarding**

#### ✅ Feature

All requests to the proxy (e.g. `/products`) are forwarded to the origin (e.g. `http://dummyjson.com/products`).

#### 📍 What to Implement

- Extract the path and query parameters from the request.
- Make a matching request to:

  ```
  <origin>/<path>?<query>
  ```

- Return:

  - Status code
  - Response body
  - Headers (copy all origin response headers)
  - Add `X-Cache: MISS` header if it's a fresh request.

---

### 3. **Response Caching**

#### ✅ Feature

If the same request is made again, return the cached response.

#### 📍 What to Implement

- Create a simple cache (can be in-memory using a map/dictionary).
- Key: Full URL including query params (e.g. `/products?page=2`)
- Value: Response object (status code, headers, body)
- If cache hit:

  - Return cached response with header: `X-Cache: HIT`

- If cache miss:

  - Forward to origin, store the response, return with `X-Cache: MISS`

---

### 4. **Cache-Control Header (Optional but Recommended)**

#### ✅ Feature

Honor basic cache control.

#### 📍 What to Implement

- If the origin response includes `Cache-Control: no-store`, don’t cache it.
- If `Cache-Control: max-age=<seconds>`, store with expiry logic.

---

### 5. **Clear Cache Command**

#### ✅ Feature

Allow user to clear the cache using:

```bash
caching-proxy --clear-cache
```

#### 📍 What to Implement

- If this command is run, load the CLI and wipe the internal cache storage.
- Show a success message like: `Cache cleared successfully.`

---

## 🗂 Example Behavior

### 🎯 When User Runs:

```bash
caching-proxy --port 3000 --origin http://dummyjson.com
```

And client sends:

```
GET http://localhost:3000/products
```

Your proxy should:

- Forward request to: `http://dummyjson.com/products`
- Cache response
- Return response with: `X-Cache: MISS`

Next time:

- Return cached response with: `X-Cache: HIT`

---

## 📁 Suggested Project Structure

```
caching-proxy/
├── main.go (or main.ts / index.js etc.)
├── cache/          # caching logic
│   └── memory.go   # or memory.ts
├── proxy/          # proxy logic
│   └── handler.go
├── cli/            # CLI parser
│   └── parser.go
├── utils/
│   └── logger.go
└── README.md
```

---

## 🧪 Optional (Bonus) Features

These are not mandatory but great if you want to impress:

- ⏱ TTL for cache entries (e.g. expire after 5 mins)
- 💾 File-based cache (save cache to disk)
- 📊 Logging of request origin, status, hit/miss
- 🧼 Admin endpoint to check and clear cache via API (e.g. `DELETE /__admin/clear-cache`)
- 🌐 Support for multiple HTTP methods (GET, POST, etc.)

---

## 🛠 Technologies You Can Use

- **Language**: Go, Node.js, Python, Rust — your choice
- **HTTP Client**: `net/http`, `axios`, `requests`, etc.
- **CLI Parsing**:

  - Go → `flag` or `cobra`
  - Node.js → `commander` or `yargs`
  - Python → `argparse` or `click`

---

## 🎯 What You Will Learn

- How to build a reverse proxy server
- Handling and parsing HTTP requests
- Implementing in-memory caching
- CLI design and argument parsing
- Header manipulation and cache validation

---
