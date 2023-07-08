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

func (r UpdateUserInfo) Rules() ValidationRules {
	return struct {
		UserID   string `validate:"required"`
		Name     string `validate:"required"`
		Email    string `validate:"required"`
		Password string `validate:"required"`
	}{
		UserID:   r.UserID,
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
