package objects

import (
	"github.com/gin-gonic/gin"
	"os"
)

func get(context *gin.Context) {
	fileDir := os.Getenv("STORAGE_ROOT")
	fileName := context.Param("value")
	context.Header("Content-Type", "application/octet-stream")
	context.Header("Content-Disposition", "attachment; filename="+fileName)
	context.Header("Content-Transfer-Encoding", "binary")
	context.File(fileDir + "/objects/" + fileName)
	return
}
