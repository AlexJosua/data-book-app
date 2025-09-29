package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        username, password, ok := c.Request.BasicAuth()
        if !ok || username != "admin" || password != "secret" {
            c.Header("WWW-Authenticate", `Basic realm="restricted"`)
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }
        c.Next()
    }
}
