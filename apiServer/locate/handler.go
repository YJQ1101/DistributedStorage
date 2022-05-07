package locate

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(engine *gin.Engine) {
	engine.GET("/locate/:value", func(context *gin.Context) {
		filename := context.Param("value")
		info := Locate(filename)
		if len(info) == 0 {
			context.String(http.StatusNotFound, "fail")
			return
		}
		context.String(http.StatusOK, info)
	})
}
