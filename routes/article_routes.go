package routes

import(
	"blogProjectGin/controllers"
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(r *gin.Engine){
	r.GET("/api/v1/List_article", controllers.List_article)
	r.POST("/api/v1/CreateArticle", controllers.CreateArticle)
	r.GET("/api/v1/ArticleDetail/:id", controllers.ArticleDetail)
	r.PUT("/api/v1/UpdateArticle/:id", controllers.UpdateArticle)
	r.DELETE("/api/v1/DeleteArticle/:id", controllers.DeleteArticle)
}




