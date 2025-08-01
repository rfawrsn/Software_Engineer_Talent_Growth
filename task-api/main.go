package main

import (
	"fmt"
	"task-api/config"
	"task-api/controllers"
	"task-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB() 

	controllers.InitTaskController()

    r := gin.Default()
    routes.SetupTaskRoutes(r)
    r.Run(":8090")
	fmt.Println("✅ MongoDB connected:", config.DB != nil)

}
