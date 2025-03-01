package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAlterEnums, downAlterEnums)
}

func upAlterEnums(ctx context.Context, tx *sql.Tx) error {

	_, err := tx.Exec(`
		ALTER TABLE users
		DROP COLUMN gender, 
		DROP COLUMN marital_status;

		ALTER TABLE users
		ADD COLUMN gender gender, 
		ADD COLUMN marital_status marital_status;
	`)
	if err != nil {
		return err
	}

	return nil

}

func downAlterEnums(ctx context.Context, tx *sql.Tx) error {

	_, err := tx.Exec(`
		ALTER TABLE users
		DROP COLUMN gender, 
		DROP COLUMN marital_status;
	`)
	if err != nil {
		return err
	}

	return nil
}
