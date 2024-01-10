package models

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
)

type Shortening struct {
	Identifier  string `json:"identifier"`
	OriginalURL string `json:"original_url"`
}

type ShortenInput struct {
	RawURL     string  `json:"raw_url" binding:"required"`
	Identifier *string `json:"identifier"`
	CreatedBy  string  `json:"created_by"`
}
