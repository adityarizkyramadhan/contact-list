package http

import (
	"net/http"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/adityarizkyramadhan/contact-list/utils"
	"github.com/gin-gonic/gin"
)

type contract struct {
	usecaseContact domain.ContactUsecase
}

func New(usecaseContact domain.ContactUsecase) *contract {
	return &contract{usecaseContact}
}

func (c *contract) Create(ctx *gin.Context) {
	var model request.ContactCreate
	if err := ctx.ShouldBindJSON(&model); err != nil {
		ctx.Error(err)
		return
	}
	if err := c.usecaseContact.Create(ctx, &model); err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, "success create contact", nil)
}

func (c *contract) FindByID(ctx *gin.Context) {
	// ambil id dari token
	id := ctx.MustGet("id").(string)
	if id == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: id is required", nil)
	}
	contactID := ctx.Param("id")
	if contactID == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: contact id is required", nil)
	}
	contact, err := c.usecaseContact.FindByID(ctx, id, contactID)
	if err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success get contact", contact)
}

func (c *contract) Update(ctx *gin.Context) {
	var model request.ContactUpdate
	if err := ctx.ShouldBindJSON(&model); err != nil {
		ctx.Error(err)
		return
	}
	if err := c.usecaseContact.Update(ctx, &model); err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success update contact", nil)
}

func (c *contract) Delete(ctx *gin.Context) {
	// ambil id dari token
	id := ctx.MustGet("id").(string)
	if id == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: id is required", nil)
	}
	contactID := ctx.Param("id")
	if contactID == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: contact id is required", nil)
	}
	if err := c.usecaseContact.Delete(ctx, id, contactID); err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success delete contact", nil)
}

func (c *contract) FindAll(ctx *gin.Context) {
	// ambil id dari token
	id := ctx.MustGet("id").(string)
	if id == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: id is required", nil)
	}
	query, err := request.BuildContactQuery(ctx.Query("name"), ctx.Query("page"), ctx.Query("limit"))
	if err != nil {
		ctx.Error(err)
		return
	}
	contacts, err := c.usecaseContact.FindAll(ctx, id, query)
	if err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success get all contact", contacts)
}
