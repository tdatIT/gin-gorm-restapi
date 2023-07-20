package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdatIT/gin-gorm-restapi/src/config"
	"github.com/tdatIT/gin-gorm-restapi/src/routers"
)

func main() {
	defer fmt.Println("Server is running on port 8080")
	config.InitConnection()
	r := gin.New()
	routers.RegisterRouters(r)
	r.Run(":8080")
}
