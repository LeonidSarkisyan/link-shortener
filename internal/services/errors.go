package services

import "errors"

var (
	BadUsernameOrPassword = errors.New("неверный логин или пароль")
)
