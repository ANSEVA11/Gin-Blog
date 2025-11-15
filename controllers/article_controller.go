package controllers

import (
	"blogProjectGin/config"
	"blogProjectGin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// list Of Articles
func List_article(context *gin.Context){

	var articles []models.Article
	config.DB.Find(&articles)
	context.JSON(http.StatusOK, articles)
}

// Create Articles 
func CreateArticle(context *gin.Context){

	var article models.Article

	err := context.ShouldBindJSON(&article)

	// Error Handling
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&article)
	context.JSON(http.StatusCreated, article)

}

// Article Detail (Get an Article)
func ArticleDetail(context *gin.Context){
	var article models.Article

	id_str := context.Param("id")
	id,_ := strconv.Atoi(id_str)

	result := config.DB.First(&article, id)
	if result.Error != nil{
		context.JSON(http.StatusNotFound, gin.H{"error": "Article Not Found..."})
		return
	}
	context.JSON(http.StatusOK, article)
}

// Update Article
func UpdateArticle(context *gin.Context){

	var article_old models.Article

	id_str := context.Param("id")
	id, _ := strconv.Atoi(id_str)

	result := config.DB.First(&article_old, id)
	if result.Error != nil{
		context.JSON(http.StatusNotFound, gin.H{"error": "Article Not found...!"})
		return
	}

	var article_new models.Article

	err := context.ShouldBindJSON(&article_new)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article_old.Title = article_new.Title
	article_old.Body = article_new.Body
	config.DB.Save(&article_old) 	// ذخیره کردن اطلاعات

	context.JSON(http.StatusOK, article_old)
}

// DELETE Article
func DeleteArticle(context *gin.Context){
	var article models.Article
	
	id_str := context.Param("id")
	id, _ := strconv.Atoi(id_str)

	result := config.DB.First(&article, id)
	if result.Error != nil{
		context.JSON(http.StatusNotFound, gin.H{"error": "Article Not found...!"})
		return
	}
	config.DB.Delete(&article)
	context.JSON(http.StatusOK, gin.H{"message": "article deleted...!"})
}

