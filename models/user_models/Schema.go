package user_models

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Age       string    `json:"age"`
	Gender    string    `json:"gender"`
	CreatedAt string    `json:"createdAt"`
	DeletedAt string    `json:"deletedAt"`
}

type UserTemplate struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Gender string    `json:"gender"`
}

type UserStructs struct {
	User         User
	UserTemplate UserTemplate
}
