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
	goose.AddMigrationContext(upCreateUser, downCreateUser)
}

func upCreateUser(ctx context.Context, tx *sql.Tx) error {

	_, err := tx.Exec(`
		CREATE TYPE gender AS ENUM ('male', 'female', 'other');
		CREATE TYPE marital_status AS ENUM ('single', 'married');
	`)

	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Migrator().CreateTable(&users.User{})

}

func downCreateUser(ctx context.Context, tx *sql.Tx) error {

	_, err := tx.Exec(`
		DROP TYPE IF EXISTS gender;
		DROP TYPE IF EXISTS marital_status;
	`)
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.Migrator().DropTable(&users.User{})
}
