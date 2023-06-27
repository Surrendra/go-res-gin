package productcontroller

import (
	"net/http"
	"fmt"

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
		"data": products,
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

func Create(c *gin.Context){
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data not found",
			"error": err.Error(),
		})
		fmt.Println(err)
		return
	}
	
	product.Code = uuid.New().String()

	// calculate age from year_born from input
	// get current year from time.Now().Year() 
	var currentYear = time.Now().Year() 
	var randomI = 20
	product.Age = currentYear - product.YearBorn + randomI
	fmt.Println(randomI)

	fmt.Println(product)
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"data": product,
		"message": "Data has been created",
	})
}

func Update(c *gin.Context){
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
			"error": err.Error(),
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
		"data": product,
		"message": "Data has been updated",
	})
}

func Delete(c *gin.Context ){
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
		"data": product,
		"message": "Data "+product.Name+" has been deleted",
	})
}
