package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

func sessionScan(rows *sql.Rows) (Session, error) {
	fields := sessionStatic()

	err := rows.Scan(
		fields.Id,
		fields.UserId,
		fields.Token,
	)
	if err != nil {
		return Session{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddSessionOut struct {
	Session Session
}

func (db *Database) AddSession(tx *sql.Tx, params Session) (*AddSessionOut, error) {
	query, args := db.gen.Add(DigitalTraceSession, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddSessionOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := sessionScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.Session = row

	return out, nil
}

// * FIND
type FindSessionOut struct {
	IsFound bool
	Session []Session
}

func (db *Database) FindSession(tx *sql.Tx, where Session) (*FindSessionOut, error) {
	query, args := db.gen.Find(DigitalTraceSession, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindSessionOut{}

	for rows.Next() {
		row, err := sessionScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addSessionScan(rows)]")
		}

		out.Session = append(out.Session, row)
	}

	if len(out.Session) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeSessionOut struct {
	IsFound bool
	Session []Session
}

func (db *Database) ChangeSession(tx *sql.Tx, set Session, where Session) (*ChangeSessionOut, error) {
	query, args := db.gen.Change(DigitalTraceSession, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeSessionOut{}

	for rows.Next() {
		row, err := sessionScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addSessionScan(rows)]")
		}

		out.Session = append(out.Session, row)
	}

	if len(out.Session) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteSessionOut struct {
	IsFound bool
	Session []Session
}

func (db *Database) DeleteSession(tx *sql.Tx, where Session) (*DeleteSessionOut, error) {
	query, args := db.gen.Delete(DigitalTraceSession, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteSessionOut{}

	for rows.Next() {
		row, err := sessionScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addSessionScan(rows)]")
		}

		out.Session = append(out.Session, row)
	}

	if len(out.Session) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
