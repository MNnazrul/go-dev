# Caching Proxy CLI Tool

A simple command-line tool that launches a caching reverse proxy server. The proxy forwards HTTP requests to a specified origin server, caches the responses, and serves cached responses for repeated requests.

---

## 🚀 Features

- **Start a proxy server via CLI**
- **Request forwarding**: Forwards all HTTP requests to the origin, preserving path and query parameters.
- **Response caching**: Caches responses in memory and serves cached responses for repeated requests.
- **Cache control**: Adds `X-Cache: HIT` or `X-Cache: MISS` headers to responses.
- **Clear cache command**: Clear the cache using a CLI flag.

---

## 🛠 Usage

### Start the Proxy

```bash
go run main.go --port <PORT> --origin <ORIGIN_URL>
```

Example:

```bash
go run main.go --port 8080 --origin https://httpbin.org
```

### Clear the Cache

```bash
go run main.go --clear-cache --port <PORT> --origin <ORIGIN_URL>
```

---

## 🗂 Project Structure

```
caching-proxy/
├── main.go
├── go.mod
├── README.md
├── project-requirements.md
├── cache/
│   └── memory.go
├── cli/
│   └── cli.go
├── proxy/
│   └── server.go
└── utils/
```

---

## 📦 Requirements

- Go 1.18 or newer

---

## 📚 Source & Inspiration

Project idea from [roadmap.sh: Caching Server Project](https://roadmap.sh/projects/caching-server)

---

## 📝 License

MIT
