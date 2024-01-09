package users

import (
	"github.com/edcnogueira/api-login/service/usersservice"
)

type Service interface {
	Create(request usersservice.CreateRequest) (usersservice.CreateResponse, error)
}

type Users struct {
	service Service
}
