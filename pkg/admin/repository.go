package admin

import (
	"context"
	"github.com/satori/go.uuid"

)

type Repository interface {
	Get(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID) (interface{}, error)
	ListProjectFromUser(ctx context.Context, ownerUuid *uuid.UUID, userUuid *uuid.UUID) (interface{}, error)
	List(ctx context.Context, ownerUuid *uuid.UUID) (interface{}, error)
	Create(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID) error
	CreateWithValue(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID, json string) error
	CreateMulti(ctx context.Context, ownerUuid *uuid.UUID, projectUuid []*uuid.UUID) error
	Update(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID) error
	Delete(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID) error
}