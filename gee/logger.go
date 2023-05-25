package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		log.Printf("[%s]%s %v", c.Method, c.Req.RequestURI, c.Params)
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))

	}
}
