package objects

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func get(context *gin.Context) {
	fileName := context.Param("value")
	stream, err := getStream(fileName)
	if err != nil {
		context.String(http.StatusNotFound, "404")
		log.Println(err)
		return
	}
	context.String(http.StatusOK, string(stream))
}
