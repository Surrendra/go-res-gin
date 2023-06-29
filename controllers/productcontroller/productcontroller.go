package productcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-res-gin/models"
	"github.com/google/uuid"
	"time"
)

func GetAllData(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been fetched",
		"data":    products,
	})
}

func FindByCode(c *gin.Context) {
	var product models.Product
	code := c.Param("code")

	if err := models.DB.Where("code = ?", code).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data not found",
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}

	authId, _ := c.Get("authId")
	var user models.User
	models.DB.Where("id = ?", authId).First(&user)

	//fmt.Println(authId)
	product.Code = uuid.New().String()
	product.CreatedUserId = user.Id

	var currentYear = time.Now().Year()
	var randomI = 20
	product.Age = currentYear - product.YearBorn + randomI

	//models.DB.Create(&product.CreatedUser

	// create user and get with user
	models.DB.Create(&product)
	models.DB.Preload("CreatedUser").First(&product)

	randomProduct := models.Product{
		Code:          uuid.New().String(),
		Name:          "Random Product",
		Description:   "Random Product Description",
		YearBorn:      1990,
		Age:           30,
		CreatedUserId: user.Id,
		Status:        true,
	}
	models.DB.Create(&randomProduct)
	c.JSON(http.StatusOK, gin.H{
		"data":           product,
		"random_product": randomProduct,
		"message":        "Data has been created",
	})
}

func Update(c *gin.Context) {
	var product models.Product
	code := c.Param("code")
	fmt.Println(code)
	if err := models.DB.Where("code = ?", code).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
		fmt.Println(err)
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data not found",
			"error":   err.Error(),
		})
		fmt.Println(err)
		return
	}

	// calculate age from year_born from input
	// get current year from time.Now().Year()
	var currentYear = time.Now().Year()
	product.Age = currentYear - product.YearBorn

	models.DB.Save(&product)

	c.JSON(http.StatusOK, gin.H{
		"data":    product,
		"message": "Data has been updated",
	})
}

func Delete(c *gin.Context) {
	var product models.Product
	code := c.Param("code")
	fmt.Println(code)
	if err := models.DB.Where("code = ?", code).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data not found",
		})
		fmt.Println(err)
		return
	}

	models.DB.Delete(&product)

	// concat string in go

	c.JSON(http.StatusOK, gin.H{
		"data":    product,
		"message": "Data " + product.Name + " has been deleted",
	})
}
