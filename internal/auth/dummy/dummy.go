package dummy

import (
	"context"
	"github.com/sirkon/weather-cacher/internal/auth"
)

var _ auth.Auth = DummyAuth{}

// DummyAuth всегда всё разрешающая авторизация
type DummyAuth struct{}

// Auth ...
func (DummyAuth) Auth(_ context.Context, token auth.Hidden) (bool, error) {
	return true, nil
}

// UserData ...
func (DummyAuth) UserData(_ context.Context, token auth.Hidden) (*auth.UserData, error) {
	data := auth.UserData{
		Name:      "Senior Developer",
		FirstName: "Senior",
		LastName:  "Developer",
		Email:     "",
	}
	return &data, nil
}
