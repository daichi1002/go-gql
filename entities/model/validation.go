package model

type ValidationRules interface{}

type Validatable interface {
	Rules() ValidationRules
}

func (r CreateUserInfo) Rules() ValidationRules {
	return struct {
		Name     string `validate:"required"`
		Email    string `validate:"required"`
		Password string `validate:"required"`
	}{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
