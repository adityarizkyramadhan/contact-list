package user

import (
	"github.com/adityarizkyramadhan/contact-list/middleware"
	delivery "github.com/adityarizkyramadhan/contact-list/user/delivery/http"
	"github.com/adityarizkyramadhan/contact-list/user/repository"
	"github.com/adityarizkyramadhan/contact-list/user/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB, router *gin.Engine) *User {
	return &User{db, router}
}

func (u *User) Init() {
	userRepo := repository.New(u.db)
	userUc := usecase.New(userRepo)
	userHttp := delivery.New(userUc)
	userRoute := u.router.Group("/user")
	{
		userRoute.POST("/register", userHttp.Register)
		userRoute.POST("/login", userHttp.Login)
		userRoute.GET("/", middleware.ValidateJWToken(), userHttp.FindByID)
		userRoute.PUT("/:id", middleware.ValidateJWToken(), userHttp.Update)
		userRoute.DELETE("/", middleware.ValidateJWToken(), userHttp.Delete)
	}
}
