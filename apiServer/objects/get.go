package objects

import (
	"DistributedStorage/apiServer/locate"
	"DistributedStorage/src/lib/es"
	"DistributedStorage/src/lib/objectstream"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func get(context *gin.Context) {
	name := context.Param("value")
	versionId := context.Query("version")

	version := 0
	var err error
	if len(versionId) != 0 {
		version, err = strconv.Atoi(versionId)
		if err != nil {
			log.Println(err)
			return
		}
	}
	meta, err := es.GetMetadata(name, version)
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "500")
		return
	}
	if meta.Hash == "" {
		context.String(http.StatusNotFound, "404")
		return
	}
	object := url.PathEscape(meta.Hash)

	stream, err := getStream(object)
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
