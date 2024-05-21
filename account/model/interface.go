package model

import (
	"context"

	"github.com/google/uuid"
)

// type UserService interface {
// 	Get(uid uuid.UUID) (*User, error)
// }

// type UserRepository interface {
// 	FindByID(uid uuid.UUID) (*User, error)
// }

// UserService defines methods the handler layer expects
// any service it interacts with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

// UserRepository defines methods the service layer expects
// any repository it interacts with to implement
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}
