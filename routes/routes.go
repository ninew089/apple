package routes

import (
	"apple/config"
	"apple/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	db := config.GetDB()
	articlesGroup := r.Group("/api/v1/articles")
	articleController := controllers.Articles{DB: db}
	custumersController := controllers.Customers{DB: db}
	{
		articlesGroup.GET("/q", articleController.FindCategory)
		articlesGroup.GET("/product", articleController.FindAll)
		articlesGroup.GET("/product/:id", articleController.FindOne)
		articlesGroup.POST("/product", articleController.Create)
		articlesGroup.GET("/customer", custumersController.FindAllCustomer)
		articlesGroup.POST("/customer", custumersController.CreateCustomer)
		
	}
	


}