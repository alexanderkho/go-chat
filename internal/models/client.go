package models

import "github.com/google/uuid"

type Client struct {
	Username string    `json:"username"`
	Id       uuid.UUID `json:"id"`
}
