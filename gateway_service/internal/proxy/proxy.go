package proxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewReverseProxy(target string) *httputil.ReverseProxy {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	return proxy
}

func ProxyRequest(c *gin.Context, target string) {
	proxy := NewReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}
