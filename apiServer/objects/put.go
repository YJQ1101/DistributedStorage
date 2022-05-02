package objects

import (
	"github.com/gin-gonic/gin"
	"log"
)

func put(context *gin.Context) {
	fileName := context.Param("value")
	c, err := storeObject(context.Request.Body, fileName)
	if err != nil {
		log.Println(err)
	}
	context.String(c, "good")
}
