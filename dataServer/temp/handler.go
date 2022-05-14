package temp

import "github.com/gin-gonic/gin"

func Handler(engine *gin.Engine) {
	engine.PUT("/temp/:uuid", put)
	engine.PATCH("/temp/:uuid", patch)
	engine.POST("/temp/:uuid", post)
	engine.DELETE("/temp/:uuid", del)
}
