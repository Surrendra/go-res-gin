package usercontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-res-gin/helpers"
	"github.com/go-res-gin/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetAllData(c *gin.Context) {
	var users []models.User
	// get user data from database where age > 20
	models.DB.Where("age > ?", 20).Find(&users)
	// models.DB.Find(&users)

	helpers.ResponseSuccess(c, users, "Sucesfullyasas fetch data")
}

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong with your input",
			"error":   err.Error(),
		})

		c.JSON(http.StatusCreated, gin.H{})
		fmt.Println(err)
		return
	}

	user.Age = time.Now().Year() - user.YearBorn
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong with your input",
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}
	user.Password = string(password)
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been created",
		"data":    user,
	})
}

func FindByCode(c *gin.Context) {
	var user models.User
	code := c.Param("code")

	if err := models.DB.Where("code = ?", code).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "Data has been fetched",
	})
}

func Update(c *gin.Context) {
	var user models.User
	code := c.Param("code")

	if err := models.DB.Where("code = ?", code).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
		fmt.Println(err)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong with your input",
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}

	user.Age = time.Now().Year() - user.YearBorn
	fmt.Println(user)
	models.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been updated",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	var user models.User
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong with your input",
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}

	// query get by email or username

	if err := models.DB.Where("email = ? or username = ?", inputUser.Username, inputUser.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username atau email tidak ditemukan",
		})
		fmt.Println(err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something wrong with your input",
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}

	// generate jwt token here
	//token, err := helpers.GenerateToken(user.Code)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"data":    user,
	})
}
