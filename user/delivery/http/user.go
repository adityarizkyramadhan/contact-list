package http

import (
	"net/http"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/adityarizkyramadhan/contact-list/utils"
	"github.com/gin-gonic/gin"
)

type user struct {
	usecaseUser domain.UserUsecase
}

func New(userUc domain.UserUsecase) *user {
	return &user{userUc}
}

func (u *user) Register(ctx *gin.Context) {
	var userReq request.UserRegister
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.Error(err)
		return
	}
	if err := u.usecaseUser.Register(ctx, &userReq); err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, "success register user", nil)
}

func (u *user) Login(ctx *gin.Context) {
	var userReq request.UserLogin
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.Error(err)
		return
	}
	token, err := u.usecaseUser.Login(ctx, &userReq)
	if err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success login user", gin.H{"token": token})
}

func (u *user) FindByID(ctx *gin.Context) {
	// ambil id dari token
	id := ctx.MustGet("id").(string)
	if id == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: id is required", nil)
	}
	user, err := u.usecaseUser.FindByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success get user", user)
}

func (u *user) Update(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	if id == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: id is required", nil)
		return
	}
	var userReq request.UserUpdate
	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		ctx.Error(err)
		return
	}
	userReq.ID = id
	if err := u.usecaseUser.Update(ctx, &userReq); err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success update user", nil)
}

func (u *user) Delete(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	if id == "" {
		utils.ResponseFail(ctx, http.StatusBadRequest, "bad request: id is required", nil)
		return
	}
	if err := u.usecaseUser.Delete(ctx, id); err != nil {
		ctx.Error(err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "success delete user", nil)
}
