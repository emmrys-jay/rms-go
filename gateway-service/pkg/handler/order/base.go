package order

import (
	"gateway-service/utility"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}
