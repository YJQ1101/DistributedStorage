package objects

import (
	"DistributedStorage/apiServer/heartbeat"
	"DistributedStorage/src/lib/es"
	"DistributedStorage/src/lib/objectstream"
	"DistributedStorage/src/lib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
)

func put(context *gin.Context) {
	hash := utils.GetHashFromHeader(context)
	if hash == "" {
		log.Println("missing object hash in digest header")
		context.String(http.StatusBadRequest, "badRequest")
		return
	}
	c, err := storeObject(context.Request.Body, url.PathEscape(hash))
	if err != nil {
		log.Println(err)
		context.String(c, "404")
		return
	}
	if c != http.StatusOK {
		context.String(c, "wrong")
		return
	}
	name := context.Param("value")
	size := utils.GetSizeFromHeader(context)
	err = es.AddVersion(name, hash, size)
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "ServerError")
		return
	}
}

func storeObject(r io.Reader, object string) (int, error) {
	stream, err := putStream(object)
	if err != nil {
		return http.StatusServiceUnavailable, err
	}
	io.Copy(stream, r)
	err = stream.Close()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewPutStream(server, object), nil
}
