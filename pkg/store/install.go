package store

import (
	"context"
)

func (s *Store) install(ctx context.Context) error {
	s.db.NewDropTable().Model((*Account)(nil)).IfExists().Exec(ctx)
	s.db.NewCreateTable().Model((*Account)(nil)).IfNotExists().Exec(ctx)
	return nil
}
