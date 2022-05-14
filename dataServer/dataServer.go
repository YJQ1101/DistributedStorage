package main

import (
	"DistributedStorage/dataServer/heartbeat"
	"DistributedStorage/dataServer/locate"
	"DistributedStorage/dataServer/objects"
	"DistributedStorage/dataServer/temp"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	objects.Handler(r)
	temp.Handler(r)
	r.Run(os.Getenv("LISTEN_ADDRESS"))
}
