package usersservice

import (
	"github.com/google/uuid"
	"time"
)

type CreateRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateResponse struct {
	User User
}

func Create(request CreateRequest) (CreateResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{User: User{
		ID:        id.String(),
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}}, nil
}
