package migrations

import (
	"context"

	"github.com/luigiescalante/bun-starter-kit/bunapp"
	"github.com/luigiescalante/bun-starter-kit/example"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		db.RegisterModel((*example.User)(nil), (*example.Org)(nil))

		fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
		return fixture.Load(ctx, bunapp.FS(), "fixture/fixture.yml")
	}, nil)
}
