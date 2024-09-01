package store

import (
	"context"
	"time"
	"github.com/uptrace/bun"
)

type Timestamp struct {
	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:",soft_delete,nullzero"`
}

var _ bun.BeforeAppendModelHook = (*Timestamp)(nil)

func (x *Timestamp) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		x.UpdatedAt = time.Now()
		x.CreatedAt = x.UpdatedAt
	case *bun.UpdateQuery:
		x.UpdatedAt = time.Now()
	}
	return nil
}

type Account struct {
	bun.BaseModel `bun:"table:hk4e_accounts"`

	ID       int64  `bun:",pk,autoincrement"`
	Email    string `bun:",notnull,unique"`
	Username string `bun:",notnull,unique"`
	Password string `bun:",nullzero"`

	LoginToken string `bun:",nullzero"`
	ComboToken string `bun:",nullzero"`

	Timestamp
}