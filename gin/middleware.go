package gin

import (
	"fmt"
	"github.com/betacraft/yaag/middleware"
	"github.com/betacraft/yaag/models"
	"github.com/betacraft/yaag/yaag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httptest"
	"strings"
)

func Document() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !yaag.IsOn() {
			return
		}
		writer := httptest.NewRecorder()
		apiCall := models.ApiCall{}
		middleware.Before(&apiCall, c.Request)
		c.Next()
		r := c.Request
		if writer.Code != 404 {
			apiCall.MethodType = c.Request.Method
			apiCall.CurrentPath = strings.Split(c.Request.RequestURI, "?")[0]
			apiCall.ResponseBody = ""
			apiCall.ResponseCode = c.Writer.Status()
			headers := map[string]string{}
			for k, v := range c.Writer.Header() {
				log.Println(k, v)
				headers[k] = strings.Join(v, " ")
			}
			apiCall.ResponseHeader = headers
			go yaag.GenerateHtml(&apiCall)
		}
	}
}
