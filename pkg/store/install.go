package store

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Store) install(ctx context.Context) error {
	if s.checkInit(ctx) {
		return nil
	}
	s.db.NewDropTable().Model((*Account)(nil)).IfExists().Exec(ctx)
	s.db.NewCreateTable().Model((*Account)(nil)).IfNotExists().Exec(ctx)
	randBytes := make([]byte, 6)
	if _, err := rand.Read(randBytes); err != nil {
		return err
	}
	base64Str := base64.RawStdEncoding.EncodeToString(randBytes)
	hashed, err := bcrypt.GenerateFromPassword([]byte(base64Str), bcrypt.DefaultCost)
	if err != nil {
		return err
	})
}
