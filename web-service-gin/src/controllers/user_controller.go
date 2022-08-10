package controllers

import (
	"fmt"
	"net/http"

	"mostafiz9900/web-service-gin/src/config"
	"mostafiz9900/web-service-gin/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := config.InitDb()
	db.AutoMigrate(&models.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)
	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, user)

}
func (repository *UserRepo) GetUsers(c *gin.Context) {
	var user []models.User
	// err := models.GetUsers(repository.Db, &user)
	result := repository.Db.Raw("SELECT id,name FROM user").Scan(&user)
	fmt.Println(result)
	fmt.Println("ekhane")
	if user != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": user})
		return
	}
	c.JSON(http.StatusOK, user)

}
func (repository *UserRepo) PostUserInfo(c *gin.Context) {
	user := models.User{Name: "mostafi", Phone: "56d78788"}
	err := models.PostUsers(repository.Db, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)

}
func GetAllUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		myMap := make(map[string]interface {
		})
		myMap["id"] = id
		myMap["name"] = "Mostafizur"
		myMap["age"] = 29

		ctx.JSON(http.StatusOK, gin.H{
			"data": myMap,
		})
	}
}
