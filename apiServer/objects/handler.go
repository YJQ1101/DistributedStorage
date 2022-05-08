package objects

import (
	"github.com/gin-gonic/gin"
)

func Handler(engine *gin.Engine) {
	engine.GET("/objects/:value", get)
	engine.PUT("/objects/:value", put)
	engine.DELETE("/objects/:value", del)
}
