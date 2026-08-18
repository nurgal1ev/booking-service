package user

import (
	"context"

	"github.com/nurgal1ev/booking-service/internal/service/user"
)

type UserHandler struct {
	userService *user.UserService
}

func NewUserHandler(userService *user.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) RegisterHandler(ctx context.Context, input *RegisterInput) (*RegisterOutput, error) {
	user := &user.User{
		FirstName: input.Body.FirstName,
		LastName:  input.Body.LastName,
		Username:  input.Body.Username,
		Email:     input.Body.Email,
		Password:  input.Body.Password,
	}

	_, err := u.userService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	resp := &RegisterOutput{}
	resp.Body.Message = "successful registration"

	return resp, nil
}

func (u *UserHandler) LoginHandler(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	token, err := u.userService.Login(ctx, input.Body.Email, input.Body.Password)
	if err != nil {
		return nil, err
	}

	resp := &LoginOutput{}
	resp.Body.AccessToken = token

	return resp, nil
}
