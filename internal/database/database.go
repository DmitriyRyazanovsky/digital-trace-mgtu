package database

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"mgtu/digital-trace/main-backend-service/internal/features"

	"github.com/pkg/errors"
)

type Database struct {
	database *sql.DB
	gen      Generator
}

type NewDatabaseIn struct {
	Database         string
	PasswordFilePath string
	UserName         string
	Url              string
	SslMode          string

	DriverName string
}

func NewDatabase(params NewDatabaseIn) (*Database, error) {
	connectLink, err := generateLink(
		params.Database,
		params.PasswordFilePath,
		params.UserName,
		params.Url,
		params.SslMode,
	)
	if err != nil {
		err = errors.Wrap(err, "[generateLink(..., ..., ..., ..., ...)]")
		return nil, err
	}

	db, err := sql.Open(params.DriverName, connectLink)
	if err != nil {
		err = errors.Wrap(err, "[generateLink(..., ..., ..., ..., ...)]")
		return nil, err
	}

	out := &Database{
		database: db,
	}

	return out, nil
}

func generateLink(database string, passwordFilePath string, userName string, url string, sslMode string) (string, error) {
	isCreated, err := features.CreateFileIfNotExists(passwordFilePath)
	if err != nil {
		err = errors.Wrap(err, "[features.CreateFileIfNotExists(passwordFilePath)]")
		return "", err
	}

	if !isCreated {
		bytes, err := features.GenerateRandomBytes(64)
		if err != nil {
			err = errors.Wrap(err, "[features.GenerateRandomBytes(32)]")
			return "", err
		}

		// Кодирование случайных байт в base64 для создания пароля
		password := base64.URLEncoding.EncodeToString(bytes)

		err = features.WriteContentInFile(passwordFilePath, password)
		if err != nil {
			err = errors.Wrap(err, "[features.WriteContentInFile(passwordFilePath, string(bytes))]")
			return "", err
		}

		fmt.Printf("GENERATED SECRET PASS FOR DATABASE: '%s'\n", password)
	}

	password, err := features.ReadContentFromFile(passwordFilePath)
	if err != nil {
		err = errors.Wrap(err, "[features.ReadContentFromFile(passwordFilePath)]")
		return "", err
	}

	out := fmt.Sprintf("%s://%s:%s@%s?sslmode=%s", database, userName, string(password), url, sslMode)

	return out, nil
}

func (d *Database) OpenTransaction() (*sql.Tx, error) {
	tx, err := d.database.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "d.database.Begin()")
	}
	return tx, nil
}

func (d *Database) RollbackTransaction(tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return errors.Wrap(err, "[tx.Rollback()]")
	}
	return nil
}

func (d *Database) CommitTransaction(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return errors.Wrap(err, "[tx.Commit()]")
	}
	return nil
}

func (d *Database) Close() error {
	err := d.database.Close()
	if err != nil {
		return errors.Wrap(err, "[d.database.Close()]")
	}

	return nil
}
