package objects

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func put(context *gin.Context) {
	filename := context.Param("value")

	file, err := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + filename)
	println(os.Getenv("STORAGE_ROOT") + "/objects/" + filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	io.Copy(file, context.Request.Body)
}
