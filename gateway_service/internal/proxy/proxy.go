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
	// Можно модифицировать запрос до проксирования
	// Например, добавить заголовки с ролью или user_id
	// userID, _ := c.Get("user_id")
	// role, _ := c.Get("role")
	// c.Request.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))
	// c.Request.Header.Set("X-User-Role", fmt.Sprintf("%v", role))

	proxy.ServeHTTP(c.Writer, c.Request)
}
