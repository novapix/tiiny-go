package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ShortenRequest struct {
	URL string `json:"url" validate:"required,url"`
	Key string `json:"key" validate:"omitempty,alphanum,max=16"`
}

func ValidateShortenRequest(req ShortenRequest) error {
	return validate.Struct(req)
}
