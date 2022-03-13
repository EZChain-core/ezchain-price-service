package admin

import (
	"context"
	"github.com/satori/go.uuid"
	d "github.com/EZChain-core/price-service/pkg/domain"
)

type UseCase interface {
	List(ctx context.Context, ownerUuid *uuid.UUID, options map[string]interface{}) ([]*d.Project, *d.Metadata, error)
	GetByUuid(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID) (*d.Project, error)
	Create(ctx context.Context, ownerUuid *uuid.UUID, project *d.Project) error
	Delete(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID) error
	AttachUser(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID, userUuids []map[string]string) error
	DetachUser(ctx context.Context, ownerUuid *uuid.UUID, projectUuid *uuid.UUID, userUuids []string) error
}