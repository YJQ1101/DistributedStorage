package objects

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func put(context *gin.Context) {
	file, err := os.Create("./objects.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	io.Copy(file, context.Request.Body)
}
