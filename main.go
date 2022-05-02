package main

import (
	"DistributedStorage/apiServer/objects"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	objects.Handler(r)
	r.Run(":8000")
}
