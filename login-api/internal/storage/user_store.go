package storage

import "login-api/internal/model"

type UserStore interface {
    GetUser(username string) (model.User, bool)
    CreateUser(user model.User) error
}