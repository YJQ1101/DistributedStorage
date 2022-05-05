package objects

import (
	"DistributedStorage/apiServer/heartbeat"
	"DistributedStorage/src/lib/objectstream"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func put(context *gin.Context) {
	fileName := context.Param("value")
	c, err := storeObject(context.Request.Body, fileName)
	if err != nil {
		log.Println(err)
	}
	context.String(c, "good")
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
