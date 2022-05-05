package main

import (
	"DistributedStorage/apiServer/objects"
	"DistributedStorage/dataServer/heartbeat"
	"DistributedStorage/dataServer/locate"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	objects.Handler(r)
	r.Run("")
}
