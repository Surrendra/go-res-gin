package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/go-res-gin/helpers"
	middlewares "github.com/go-res-gin/middlewares"
	"github.com/go-res-gin/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserHandler struct {
	UuidHelper       helpers.UuidHelper
	filesystemHelper helpers.FileSystemHelper
	JwtMiddleware    middlewares.JwtMiddleware
}

var resHelper = helpers.NewResHelper()

func NewUserHandler(uuid helpers.UuidHelper, filesystem helpers.FileSystemHelper, jwtMiddleware middlewares.JwtMiddleware) *UserHandler {
	return &UserHandler{
		UuidHelper:       uuid,
		filesystemHelper: filesystem,
		JwtMiddleware:    jwtMiddleware,
	}
}

func (h UserHandler) GetAllUser(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	//fmt.Println(authName)
	authName, _ := c.Get("authName")

	resHelper.ResponseSuccess(c, users, "Success get data from user "+authName.(string))
}

func (h UserHandler) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		resHelper.ResponseFailed(c, err, "Something wrong when binding the json")
		return
	}
	user.Code = h.UuidHelper.GenerateUuid()
	user.Age = time.Now().Year() - user.YearBorn

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		resHelper.ResponseFailed(c, err, "Failed")
		return
	}
	user.Password = string(password)
	models.DB.Create(&user)
	//user.Token = h.UuidHelper.GenerateTokenForUser()
	signToken, err := h.JwtMiddleware.GenerateJWTToken(user)
	if err != nil {
		fmt.Println("failed generate jwt token")
		resHelper.ResponseFailed(c, err, "Failed generate jwt token")
		return
	}
	user.Token = signToken

	// save token to client cookies
	c.SetCookie("token", signToken, 3600, "/", "localhost", false, true)

	resHelper.ResponseSuccess(c, user, "Successfully find data")
}

func (h UserHandler) FindByCode(c *gin.Context) {
	code := c.Param("code")
	var user models.User
	if err := models.DB.Where("code=?", code).First(&user).Error; err != nil {
		fmt.Println(err)
		fmt.Println("user with code: " + code + " is not found")
		resHelper.ResponseFailed(c, err, "Failed")
		return
	}
	resHelper.ResponseSuccess(c, user, "Successfully find data")
	return
}

func (h UserHandler) UpdateProfilePicture(c *gin.Context) {
	photo, errPhoto := c.FormFile("photo")
	filename := c.Request.FormValue("filename")
	code := c.Param("code")

	fmt.Println("code :" + code + " filename:" + filename)

	if errPhoto != nil {
		fmt.Println(errPhoto)
		resHelper.ResponseFailed(c, errPhoto, "Failed get file from request")
		return
	}
	//errSaveUploadFile := c.SaveUploadedFile(photo, "uploads/"+photo.Filename)
	//if errSaveUploadFile != nil {
	//	fmt.Println(errSaveUploadFile)
	//	resHelper.ResponseFailed(c, errSaveUploadFile, "Failed upload file to server")
	//	return
	//}

	filenamePath, err := h.filesystemHelper.UploadFile(c, filename, "uploads/"+photo.Filename)
	if err != nil {
		fmt.Println("Errror uploading file")
		resHelper.ResponseFailed(c, err, "Failed")
		return
	}
	resHelper.ResponseSuccess(c, filenamePath, "Successfully save file")
}

func (h UserHandler) CheckPayload(c *gin.Context) {
	var Payload struct {
		Name         string   `json:"name"`
		Email        string   `json:"email"`
		Phone        int64    `json:"phone"`
		FirstNumber  int64    `json:"first_number"`
		SecondNumber int64    `json:"second_number"`
		Address      []string `json:"address"`
	}
	if err := c.ShouldBindJSON(&Payload); err != nil {
		fmt.Println(err)
		resHelper.ResponseFailed(c, err, "Failed")
		return
	}
	var Result = Payload.FirstNumber + Payload.SecondNumber
	fmt.Println(Result)
}

func (h UserHandler) CheckPayloadWithObject(c *gin.Context) {
	type Payload struct {
		CustomerName string `json:"customer_name"`
		PackageID    int    `json:"package_id"`
		PackageName  string `json:"package_name"`
		Place        string `json:"place"`
		Items        []struct {
			ItemID   int    `json:"item_id"`
			ItemName string `json:"item_name"`
			Qty      int    `json:"qty"`
			Price    int    `json:"price"`
			Subtotal int
			Optional struct {
				WithSpoon bool `json:"with_spoon"`
			} `json:"optional"`
		} `json:"items"`
	}

	var payload Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		fmt.Println(err)
		resHelper.ResponseFailed(c, err, "Failed binding json request")
		return
	}
	// foreach items in payload
	var subTotal int
	var totalTransaction int = 0
	for _, item := range payload.Items {
		subTotal = item.Qty * item.Price
		totalTransaction += subTotal
		//fmt.Println(item.ItemName)
		//fmt.Println(item.Qty)
		//fmt.Println(item.Price)
		//fmt.Println(item.Optional.WithSpoon)
	}
	authName, _ := c.Get("authName")
	authId, _ := c.Get("authId")
	//auth, _ := c.Get("auth")
	// print message to console
	fmt.Printf("var1 = %T\n", authId)
	fmt.Printf("var1 = %T\n", authName)

	fmt.Println("Auth Name: ", authName)
	//fmt.Println("Auth ID: " + authId)
	resHelper.ResponseSuccess(c, payload, "Success get data from payload")

}
