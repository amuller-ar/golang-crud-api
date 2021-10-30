package middleware

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func PanicRecovery() gin.HandlerFunc {

	logger := log.New(os.Stdout, "", log.LstdFlags)

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				r, dumError := httputil.DumpRequest(c.Request, true)

				request := string(r)

				if dumError != nil {
					request = dumError.Error()
				}

				logger.Printf("[Recovery] panic recovered: \n%s\n%s\n%s", request, err, debug.Stack())

				if !c.IsAborted() {
					c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error (recovered)")
				}
			}
		}()
		c.Next()
	}
}
