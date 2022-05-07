package main

import (
	"DistributedStorage/apiServer/heartbeat"
	"DistributedStorage/apiServer/locate"
	"DistributedStorage/apiServer/objects"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	go heartbeat.ListenHeartbeat()
	objects.Handler(r)
	locate.Handler(r)
	r.Run(os.Getenv("LISTEN_ADDRESS"))
}
