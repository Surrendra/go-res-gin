package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-res-gin/controllers/productcontroller"
	"github.com/go-res-gin/controllers/usercontroller"
	"github.com/go-res-gin/handlers"
	"github.com/go-res-gin/helpers"
	"github.com/go-res-gin/middlewares"
	"github.com/go-res-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	v1 := r.Group("/api")
	{
		product := v1.Group("/product")
		{
			product.GET("get_data", productcontroller.GetAllData)
			product.GET("find/:code", productcontroller.FindByCode)
			product.POST("create", productcontroller.Create)
			product.PUT("update/:code", productcontroller.Update)
			product.DELETE("delete/:code", productcontroller.Delete)
		}

		// besok lanjut pasang middleware

		user := v1.Group("/user")
		{
			UuidHelper := helpers.NewUuidHelper()
			filesystemHelper := helpers.NewFileSystemHelper()
			jwtMiddleware := middlewares.NewJwtMiddleware()
			userHandler := handlers.NewUserHandler(UuidHelper, filesystemHelper, jwtMiddleware)
			user.GET("get_data", userHandler.GetAllUser)
			user.GET("find/:code", userHandler.FindByCode)
			user.POST("create", userHandler.Create)
			user.PUT("update/:code", usercontroller.Update)
			user.POST("login", usercontroller.Login)
			user.PUT("update_profile_picture/:code", userHandler.UpdateProfilePicture)
			user.POST("check_payload", userHandler.CheckPayload)
		}
	}
	r.Run()
}
