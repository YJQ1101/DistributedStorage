package objects

import (
	"DistributedStorage/src/lib/es"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func del(context *gin.Context) {
	name := context.Param("value")
	version, err := es.SearchLatestVersion(name)
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "500")
		return
	}
	err = es.PutMetadata(name, version.Version+1, 0, "")
	if err != nil {
		log.Println(err)
		context.String(http.StatusInternalServerError, "500")
		return
	}
	context.String(http.StatusOK, "del ok")
}
