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
	jwtMiddleware := middlewares.NewJwtMiddleware()
	v1 := r.Group("/api")
	{
		transaction := v1.Group("transaction")
		{

			transactionHandler := handlers.NewTransactionHandler(jwtMiddleware)
			transaction.Use(jwtMiddleware.AuthMiddleware)
			transaction.GET("/get_data", transactionHandler.GetData)
			transaction.POST("create", transactionHandler.Create)
		}
		product := v1.Group("/product")
		{
			product.Use(jwtMiddleware.AuthMiddleware)
			product.GET("get_data", productcontroller.GetAllData)
			product.GET("find/:code", productcontroller.FindByCode)
			product.POST("create", productcontroller.Create)
			product.PUT("update/:code", productcontroller.Update)
			product.DELETE("delete/:code", productcontroller.Delete)

		}

		// besok lanjut pasang middleware

		// group by middleware

		user := v1.Group("/user")
		{
			UuidHelper := helpers.NewUuidHelper()
			filesystemHelper := helpers.NewFileSystemHelper()
			userHandler := handlers.NewUserHandler(UuidHelper, filesystemHelper, jwtMiddleware)
			user.POST("create", userHandler.Create)
			user.Use(jwtMiddleware.AuthMiddleware)
			user.GET("get_data", jwtMiddleware.AuthMiddleware, userHandler.GetAllUser)
			user.GET("find/:code", userHandler.FindByCode)
			user.PUT("update/:code", usercontroller.Update)
			user.POST("login", usercontroller.Login)
			user.PUT("update_profile_picture/:code", userHandler.UpdateProfilePicture)
			user.POST("check_payload", userHandler.CheckPayload)
			user.POST("check_payload_with_object", userHandler.CheckPayloadWithObject)
		}

	}
	r.Run()
}
