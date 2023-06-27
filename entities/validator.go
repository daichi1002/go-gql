package entities

import (
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/go-playground/validator/v10"
)

func Validate(target model.Validatable) error {
	v := validator.New()

	return v.Struct(target.Rules())
}
