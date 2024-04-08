package usecase

import (
	"context"
	"fmt"

	"github.com/adityarizkyramadhan/contact-list/domain"
	"github.com/adityarizkyramadhan/contact-list/middleware"
	"github.com/adityarizkyramadhan/contact-list/request"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	repoUser domain.UserRepository
}

func New(repoUser domain.UserRepository) domain.UserUsecase {
	return &user{repoUser}
}

func (u *user) Register(ctx context.Context, user *request.UserRegister) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	userDomain := &domain.User{
		Username: user.Username,
		Password: string(hashPassword),
	}
	if err := u.repoUser.Store(ctx, userDomain); err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	return nil
}

func (u *user) Login(ctx context.Context, user *request.UserLogin) (string, error) {
	if err := user.Validate(); err != nil {
		return "", fmt.Errorf("bad request: %v", err.Error())
	}
	userDomain, err := u.repoUser.FindByUsername(ctx, user.Username)
	if err != nil {
		return "", fmt.Errorf("internal server error: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(user.Password)); err != nil {
		return "", fmt.Errorf("bad request: %v", err.Error())
	}
	token, err := middleware.GenerateJWToken(userDomain.ID.String())
	if err != nil {
		return "", fmt.Errorf("internal server error: %v", err.Error())
	}
	return token, nil
}

func (u *user) FindByID(ctx context.Context, id string) (*domain.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("bad request: %v", err.Error())
	}
	user, err := u.repoUser.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("internal server error: %v", err.Error())
	}
	return user, nil
}

func (u *user) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	user, err := u.repoUser.FindByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("internal server error: %v", err.Error())
	}
	return user, nil
}

func (u *user) Update(ctx context.Context, user *request.UserUpdate) error {
	if err := user.Validate(); err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	idUser, err := uuid.Parse(user.ID)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	userDomain, err := u.repoUser.FindByID(ctx, idUser)
	if err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	if user.Username != "" {
		isExist, err := u.repoUser.FindByUsername(ctx, user.Username)
		if err != nil && err.Error() != "record not found" {
			return fmt.Errorf("internal server error: %v", err.Error())
		}
		if isExist != nil {
			return fmt.Errorf("bad request: username already exist")
		}
		userDomain.Username = user.Username
	}
	if user.Password != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("internal server error: %v", err.Error())
		}
		userDomain.Password = string(hashPassword)
	}
	if err := u.repoUser.Update(ctx, userDomain); err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	return nil
}

func (u *user) Delete(ctx context.Context, id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("bad request: %v", err.Error())
	}
	if err := u.repoUser.Delete(ctx, userID); err != nil {
		return fmt.Errorf("internal server error: %v", err.Error())
	}
	return nil
}
