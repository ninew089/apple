package controllers

import (
	"apple/models"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	pq "github.com/lib/pq"
)

type Articles struct {
	DB *gorm.DB
}
type Customers struct {
	DB *gorm.DB
}



type createArticleForm struct {
	Title   string                `form:"title" binding:"required"`
	Category string 							`form:"category" binding:"required"`
	Body    string                `form:"body" binding:"required"`
	Price 	string                `form:"price" binding:"required"`
	Image   *multipart.FileHeader `form:"image" binding:"required"`
}
type createCustomerForm struct {

	Name   string `form:"name" binding:"required"`
	Tel string `form:"tel" binding:"required"`
	Address string `form:"address" binding:"required"`
	Paid int `form:"paid" binding:"required"`
	ProductId  pq.Int64Array  `form:"productId" binding:"required"`
}


type articleResponse struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Category string `json:"category"`
	Price string `json:"price"`
	Body    string `json:"body"`
	Image   string `json:"image"`

}

type customerRespone struct {
	gorm.Model
	ID      uint   `json:"id"`
	Name   string `json:"name"`
	Tel string `json:"tel"`
	Address string `json:"address"`
	Paid int `json:"paid"`
	ProductId  pq.Int64Array  `json:"productId"`
}

func (a *Articles) FindAll(ctx *gin.Context) {
	var articles []models.Article

	a.DB.Find(&articles)

	var serializedArticles []articleResponse
	copier.Copy(&serializedArticles, &articles)
	ctx.JSON(http.StatusOK, gin.H{"articles": serializedArticles})
}
func (a *Articles) FindCategory(ctx *gin.Context){
	var articles []models.Article
	category := ctx.Query("category")


	 a.DB.Find(&articles, "category = ? ",category)

	var serializedArticles []articleResponse
	copier.Copy(&serializedArticles, &articles)
	ctx.JSON(http.StatusOK, gin.H{"articles": serializedArticles})
}
func (c *Customers) FindAllCustomer(ctx *gin.Context) {
	var customers []models.Customer

	c.DB.Find(&customers)

	var serializedArticles []customerRespone
	copier.Copy(&serializedArticles, &customers)
	ctx.JSON(http.StatusOK, gin.H{"customers": serializedArticles})
}



func (a *Articles) FindOne(ctx *gin.Context) {
	article, err := a.findArticleByID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	serializedArticle := articleResponse{}
	copier.Copy(&serializedArticle, &article)
	ctx.JSON(http.StatusOK, gin.H{"article": serializedArticle})
}



func (a *Articles) findArticleByID(ctx *gin.Context) (*models.Article, error) {
	var article models.Article
	id := ctx.Param("id")

	if err := a.DB.First(&article, id).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (a *Articles) Create(ctx *gin.Context) {
	var form createArticleForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var article models.Article
	copier.Copy(&article, &form)

	if err := a.DB.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	a.setArticleImage(ctx, &article)
	serializedArticle := articleResponse{}
	copier.Copy(&serializedArticle, &article)

	ctx.JSON(http.StatusCreated, gin.H{"article": serializedArticle})
}

func (a *Articles) setArticleImage(ctx *gin.Context, article *models.Article) error {
	file, err := ctx.FormFile("image")
	if err != nil || file == nil {
		return err
	}

	if article.Image != "" {
		article.Image = strings.Replace(article.Image, os.Getenv("HOST"), "", 1)
		pwd, _ := os.Getwd()
		os.Remove(pwd + article.Image)
	}

	path := "uploads/articles/" + strconv.Itoa(int(article.ID))
	os.MkdirAll(path, 0755)
	filename := path + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		return err
	}

	article.Image = os.Getenv("HOST")  + filename
	a.DB.Save(article)

	return nil
}

func (c *Customers) CreateCustomer(ctx *gin.Context) {
	var form createCustomerForm
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customer
	copier.Copy(&customer, &form)

	if err := c.DB.Create(&customer).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	serializedArticle := customerRespone{}
	copier.Copy(&serializedArticle, &customer)

	ctx.JSON(http.StatusCreated, gin.H{"article": serializedArticle})
}
