package storage

import "login-api/internal/model"

type UserStore interface {
    GetUser(email string) (model.User, bool)
    CreateUser(user model.User) error
    UpdateUser(oldEmail string, user model.User) error
}