package objects

import (
	"DistributedStorage/dataServer/locate"
	"DistributedStorage/src/lib/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
)

func get(context *gin.Context) {
	file := getFile(context.Param("uuid"))
	if file == "" {
		context.String(http.StatusNotFound, "404")
		return
	}
	context.File(file)
}

func getFile(hash string) string {
	file := os.Getenv("STORAGE_ROOT") + "/objects/" + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}
