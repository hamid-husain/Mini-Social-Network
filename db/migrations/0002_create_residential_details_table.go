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
	goose.AddMigrationContext(upResidentialDetails, downResidentialDetails)
}

func upResidentialDetails(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().CreateTable(&users.ResidentialDetail{})
}

func downResidentialDetails(ctx context.Context, tx *sql.Tx) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.Migrator().DropTable(&users.ResidentialDetail{})
}
