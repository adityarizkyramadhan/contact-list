package contact

import (
	delivery "github.com/adityarizkyramadhan/contact-list/contact/delivery/http"
	"github.com/adityarizkyramadhan/contact-list/contact/repository"
	"github.com/adityarizkyramadhan/contact-list/contact/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Contact struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB, router *gin.Engine) *Contact {
	return &Contact{db, router}
}

func (c *Contact) Init() {
	contactRepo := repository.New(c.db)
	contactUc := usecase.New(contactRepo)
	contactHttp := delivery.New(contactUc)
	contactRoute := c.router.Group("/contact")
	{
		contactRoute.POST("/", contactHttp.Create)
		contactRoute.GET("/:id", contactHttp.FindByID)
		contactRoute.PUT("/:id", contactHttp.Update)
		contactRoute.DELETE("/:id", contactHttp.Delete)
	}
}
