package repositories

import "errors"

var (
	UserAlreadyExists = errors.New("пользователь с таким username уже существует")
)
