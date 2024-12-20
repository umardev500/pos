package pkg

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormInstance struct {
	db *gorm.DB
}

var (
	gormInstance *GormInstance
	gormOnce     sync.Once
)

func NewGorm() *GormInstance {
	gormOnce.Do(func() {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		dbName := os.Getenv("DB_NAME")

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to database")
		}

		gormInstance = &GormInstance{
			db: db,
		}
	})

	return gormInstance
}

func (g *GormInstance) WithTransaction(ctx context.Context, f func(context.Context) error) error {
	err := g.db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, TransactionKey, tx)
		if err := f(ctx); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (g *GormInstance) GetConn(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TransactionKey).(*gorm.DB); ok {
		return g.applyDeletedFlag(tx, ctx)
	}

	return g.applyDeletedFlag(g.db.WithContext(ctx), ctx)
}

func (g *GormInstance) applyDeletedFlag(conn *gorm.DB, ctx context.Context) *gorm.DB {
	if unscoped, ok := ctx.Value(UnscopedKey).(bool); ok && unscoped {
		return conn.WithContext(ctx).Unscoped()
	}

	return conn.WithContext(ctx)
}
