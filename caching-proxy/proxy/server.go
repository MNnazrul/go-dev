package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/MNnazrul/caching-proxy/cache"
)

func StartProxy(port int, origin string, memCache *cache.MemoryCache) error {
	originURL, err := url.Parse(origin)
	if err != nil {
		return fmt.Errorf("invalid origin URL: %w", err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		cacheKey := r.URL.Path + "?" + r.URL.RawQuery

		if cached, ok := memCache.Get(cacheKey); ok {
			for k, v := range cached.Header {
				for _, vv := range v {
					w.Header().Add(k, vv)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(cached.StatusCode)
			w.Write(cached.Body)
			return
		}

		proxyURL := *originURL
		proxyURL.Path = r.URL.Path
		proxyURL.RawQuery = r.URL.RawQuery

		req, err := http.NewRequest(r.Method, proxyURL.String(), r.Body)
		if err != nil {
			http.Error(w, "Falied to create request", http.StatusInternalServerError)
			return
		}
		req.Header = r.Header.Clone()

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "Failed to reach origin", http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Falied to read response", http.StatusInternalServerError)
			return
		}

		memCache.Set(cacheKey, cache.CachedResponse{
			StatusCode: resp.StatusCode,
			Header:     resp.Header.Clone(),
			Body:       body,
		})

		for k, v := range resp.Header {
			for _, vv := range v {
				w.Header().Add(k, vv)
			}
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}

	http.HandleFunc("/", handler)
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Proxy server on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}
