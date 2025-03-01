package migrations

import (
	"context"
	"database/sql"
	"mini-social-network/models/baseModel"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upCreateOfficeDetailsTable, downCreateOfficeDetailsTable)
}

func upCreateOfficeDetailsTable(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().CreateTable(&baseModel.OfficeDetail{})
}

func downCreateOfficeDetailsTable(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().DropTable(&baseModel.OfficeDetail{})
}
