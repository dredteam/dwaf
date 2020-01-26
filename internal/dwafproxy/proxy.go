package dwafproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance *httputil.ReverseProxy
	proxyUrl *url.URL
)

func New(u *url.URL) *httputil.ReverseProxy {
	once.Do(func() {
		proxyUrl = u
		instance = newSingleHostReverseProxy()
	})
	return instance
}

func newSingleHostReverseProxy() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: director, ErrorHandler: defaultErrorHandler}
}

func director(req *http.Request) {
	target := proxyUrl
	targetQuery := target.RawQuery
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
	if targetQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	}
	if _, ok := req.Header["User-Agent"]; !ok {
		// explicitly disable User-Agent so it's not set to default value
		req.Header.Set("User-Agent", "")
	}
}

func defaultErrorHandler(rw http.ResponseWriter, req *http.Request, err error) {
	status := http.StatusBadGateway
	rw.WriteHeader(status)
	_, _ = rw.Write([]byte(http.StatusText(status) + "\n"))
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
