package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func StartProxy(port int, origin string) error {
	originURL, err := url.Parse(origin)
	if err != nil {
		return fmt.Errorf("invalid origin URL: %w", err)
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
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
