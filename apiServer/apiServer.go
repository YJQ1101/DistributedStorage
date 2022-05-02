package main

import (
	"DistributedStorage/apiServer/heartbeat"
	"DistributedStorage/apiServer/locate"
	"DistributedStorage/apiServer/objects"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	go heartbeat.ListenHeartbeat()
	objects.Handler(r)
	locate.Handler(r)
	r.Run(":8000")
}
