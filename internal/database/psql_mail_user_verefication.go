package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

func mailUserVereficationScan(rows *sql.Rows) (MailUserVerefication, error) {
	fields := mailUserVereficationStatic()

	err := rows.Scan(
		fields.Id,
		fields.CreatedAt,
		fields.Token,
		fields.Email,
		fields.Login,
		fields.Name,
		fields.Surname,
		fields.Password,
	)
	if err != nil {
		return MailUserVerefication{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddMailUserVereficationOut struct {
	MailUserVerefication MailUserVerefication
}

func (db *Database) AddMailUserVerefication(tx *sql.Tx, params MailUserVerefication) (*AddMailUserVereficationOut, error) {
	query, args := db.gen.Add(DigitalTraceMailUserVerefication, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddMailUserVereficationOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := mailUserVereficationScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.MailUserVerefication = row

	return out, nil
}

// * FIND
type FindMailUserVereficationOut struct {
	IsFound              bool
	MailUserVerefication []MailUserVerefication
}

func (db *Database) FindMailUserVerefication(tx *sql.Tx, where MailUserVerefication) (*FindMailUserVereficationOut, error) {
	query, args := db.gen.Find(DigitalTraceMailUserVerefication, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindMailUserVereficationOut{}

	for rows.Next() {
		row, err := mailUserVereficationScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[mailUserVereficationScan(rows)]")
		}

		out.MailUserVerefication = append(out.MailUserVerefication, row)
	}

	if len(out.MailUserVerefication) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeMailUserVereficationOut struct {
	IsFound              bool
	MailUserVerefication []MailUserVerefication
}

func (db *Database) ChangeMailUserVerefication(tx *sql.Tx, set MailUserVerefication, where MailUserVerefication) (*ChangeMailUserVereficationOut, error) {
	query, args := db.gen.Change(DigitalTraceMailUserVerefication, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeMailUserVereficationOut{}

	for rows.Next() {
		row, err := mailUserVereficationScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[mailUserVereficationScan(rows)]")
		}

		out.MailUserVerefication = append(out.MailUserVerefication, row)
	}

	if len(out.MailUserVerefication) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteMailUserVereficationOut struct {
	IsFound              bool
	MailUserVerefication []MailUserVerefication
}

func (db *Database) DeleteMailUserVerefication(tx *sql.Tx, where MailUserVerefication) (*DeleteMailUserVereficationOut, error) {
	query, args := db.gen.Delete(DigitalTraceMailUserVerefication, where)
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteMailUserVereficationOut{}

	for rows.Next() {
		row, err := mailUserVereficationScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addMailUserVereficationScan(rows)]")
		}

		out.MailUserVerefication = append(out.MailUserVerefication, row)
	}

	if len(out.MailUserVerefication) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
