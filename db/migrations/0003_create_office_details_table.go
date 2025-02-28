package migrations

import (
	"context"
	"database/sql"
	"project/internal/models/users"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upOfficeDetails, downOfficeDetails)
}

func upOfficeDetails(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().CreateTable(&users.OfficeDetail{})
}

func downOfficeDetails(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().DropTable(&users.OfficeDetail{})
}
