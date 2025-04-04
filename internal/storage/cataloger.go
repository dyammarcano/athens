package storage

import (
	"context"

	"github.com/gomods/athens/internal/paths"
)

// Cataloger is the interface that lists all the modules and version contained in the storage.
type Cataloger interface {
	// Catalog gets all the modules / versions.
	Catalog(ctx context.Context, token string, pageSize int) ([]paths.AllPathParams, string, error)
}
