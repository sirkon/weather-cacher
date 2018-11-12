package auth

import (
	"context"
)

// UserData данные по пользователю получаемые от авторизационного сервиса
type UserData struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Hidden скрытое значение
type Hidden string

func (Hidden) String() string {
	return "★★★★★"
}

// Auth методы авторизации пользователя
type Auth interface {
	Auth(ctx context.Context, token Hidden) (bool, error)
	UserData(ctx context.Context, token Hidden) (*UserData, error)
}
