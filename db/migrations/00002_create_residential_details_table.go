package migrations

import (
	"context"
	"database/sql"
	"mini-social-network/db/base_model"

	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	goose.AddMigrationContext(upCreateResidentialDetailsTable, downCreateResidentialDetailsTable)
}

func upCreateResidentialDetailsTable(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(&base_model.ResidentialDetail{})
}

func downCreateResidentialDetailsTable(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().DropTable(&base_model.ResidentialDetail{})
}
