package objects

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func get(context *gin.Context) {
	fileDir := "."
	fileName := "objects.txt"
	stream, err := getStream(object)
	if err != nil {
		context.String(http.StatusNotFound, "404")
		log.Println(err)
		return
	}
	context.Header("Content-Type", "application/octet-stream")
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	context.Header("Content-Transfer-Encoding", "binary")
	context.File(fileDir + "/" + fileName)
	return
}
