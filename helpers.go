package main

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyRequest(c *gin.Context, method string, targetURL *url.URL) {
	req, err := http.NewRequest(method, targetURL.String(), c.Request.Body)
	if err != nil {
		c.String(500, "Error creating proxy request: %s", err)
		return
	}

	for key, values := range c.Request.Header {
		for _, value := range values {
			if key == "Authorization" {
				continue
			}
			req.Header.Add(key, value)
		}
	}

	req.ContentLength = c.Request.ContentLength

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.String(500, "Error performing proxy request: %s", err)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	c.Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)
}
