package main

import(
	_"fmt"
	"blogProjectGin/config"
	"blogProjectGin/models"
	"blogProjectGin/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Article{})

	r := gin.Default()

	routes.ArticleRoutes(r)

	r.Run("localhost:8080")


}


