package store

import (
	"context"
	"database/sql"
	"github.com/mumingluan/axuanhttp/pkg/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Store struct {
	db         *bun.DB
	account    *AccountStore
}


func (s *Store) init() {
	sqlite, err := sql.Open(sqliteshim.ShimName, s.config.Database.DSN)
	if err != nil {
		panic(err)
	}
	sqlite.SetMaxOpenConns(1)
	s.db = bun.NewDB(sqlite, sqlitedialect.New())
	s.account = &AccountStore{db: s.db}
	if err := s.install(context.Background()); err != nil {
		panic(err)
	}
}

func (s *Store) Account() *AccountStore       { return s.account }
