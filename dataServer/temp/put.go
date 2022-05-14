package temp

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func put(context *gin.Context) {
	uuid := context.Param("uuid")
	tempinfo, err := readFromFile(uuid)
	if err != nil {
		log.Println(err)
		return
	}
	infoFile := os.Getenv("STORAGE_ROOT") + "/temp/" + uuid
	datFile := infoFile + ".dat"
	file, err := os.Open(datFile)

	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}
	actual := info.Size()
	os.Remove(infoFile)
	if actual != tempinfo.Size {
		os.Remove(datFile)
		log.Println("actual size mismatch, expect", tempinfo.Size, "actual", actual)
		context.String(http.StatusInternalServerError, "serverError")
		return
	}
	commitTempObject(datFile, tempinfo)
}
