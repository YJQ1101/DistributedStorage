package versions

import (
	"DistributedStorage/src/lib/es"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func Handler(engine *gin.Engine) {
	engine.GET("/versions/:version", func(context *gin.Context) {
		name := context.Param("version")
		from := 0
		size := 1000
		for {
			metas, err := es.SearchAllVersions(name, from, size)
			if err != nil {
				log.Println(err)
				return
			}
			for i := range metas {
				b, _ := json.Marshal(metas[i])
				// TODO
				context.Data(200, "content-length", b)
			}
			if len(metas) != size {
				return
			}
			from += size
		}
	})
}
