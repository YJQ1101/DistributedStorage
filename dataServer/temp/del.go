package temp

import (
	"github.com/gin-gonic/gin"
	"os"
)

func del(context *gin.Context) {
	uuid := context.Param("uuid")
	infoFile := os.Getenv("STORAGE_ROOT") + "/temp/" + uuid
	datFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(datFile)
}
