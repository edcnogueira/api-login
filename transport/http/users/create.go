package users

import (
	"github.com/edcnogueira/api-login/service/usersservice"
	"github.com/gofiber/fiber/v2"
)

type UserCreateRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateRequest struct {
	Body UserCreateRequestBody
}

type UserCreateResponseBody struct {
	User User `json:"user"`
}

type UserCreateResponse struct {
	Body UserCreateResponseBody
}

func (u *Users) create(ctx *fiber.Ctx) error {
	var request UserCreateRequest
	if err := ctx.BodyParser(&request.Body); err != nil {
		return err
	}

	createResponse, err := u.service.Create(usersservice.CreateRequest{
		Name:     request.Body.Name,
		Email:    request.Body.Email,
		Password: request.Body.Password,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(UserCreateResponseBody{User: toUser(&createResponse.User)})
}

func toUser(user *usersservice.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
