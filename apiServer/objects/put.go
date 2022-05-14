package objects

import (
	"DistributedStorage/apiServer/heartbeat"
	"DistributedStorage/apiServer/locate"
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
	size := utils.GetSizeFromHeader(context)
	c, err := storeObject(context.Request.Body, url.PathEscape(hash), size)
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
	err = es.AddVersion(name, hash, size)
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "ServerError")
		return
	}
}

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	stream, err := putStream(hash, size)
	if err != nil {
		return http.StatusServiceUnavailable, err
	}

	reader := io.TeeReader(r, stream)
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}

func putStream(hash string, size int64) (*objectstream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}
	return objectstream.NewTempPutStream(server, hash, size)
}
