package objects

import (
	"DistributedStorage/apiServer/locate"
	"DistributedStorage/src/lib/objectstream"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	context.String(http.StatusOK, buf.String())
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
