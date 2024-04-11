package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * Attempt

func attemptScan(rows *sql.Rows) (Attempt, error) {
	fields := attemptStatic()

	err := rows.Scan(
		fields.Id,
		fields.UserId,
		fields.TestId,
		fields.StatusId,
	)
	if err != nil {
		return Attempt{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddAttemptOut struct {
	Attempt Attempt
}

func (db *Database) AddAttempt(tx *sql.Tx, params Attempt) (*AddAttemptOut, error) {
	query, args := db.gen.Add(DigitalTraceAttempt, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddAttemptOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := attemptScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.Attempt = row

	return out, nil
}

// * FIND
type FindAttemptOut struct {
	IsFound bool
	Attempt []Attempt
}

func (db *Database) FindAttempt(tx *sql.Tx, where Attempt) (*FindAttemptOut, error) {
	query, args := db.gen.Find(DigitalTraceAttempt, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindAttemptOut{}

	for rows.Next() {
		row, err := attemptScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addAttemptScan(rows)]")
		}

		out.Attempt = append(out.Attempt, row)
	}

	if len(out.Attempt) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeAttemptOut struct {
	IsFound bool
	Attempt []Attempt
}

func (db *Database) ChangeAttempt(tx *sql.Tx, set Attempt, where Attempt) (*ChangeAttemptOut, error) {
	query, args := db.gen.Change(DigitalTraceAttempt, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeAttemptOut{}

	for rows.Next() {
		row, err := attemptScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addAttemptScan(rows)]")
		}

		out.Attempt = append(out.Attempt, row)
	}

	if len(out.Attempt) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteAttemptOut struct {
	IsFound bool
	Attempt []Attempt
}

func (db *Database) DeleteAttempt(tx *sql.Tx, where Attempt) (*DeleteAttemptOut, error) {
	query, args := db.gen.Delete(DigitalTraceAttempt, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteAttemptOut{}

	for rows.Next() {
		row, err := attemptScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addAttemptScan(rows)]")
		}

		out.Attempt = append(out.Attempt, row)
	}

	if len(out.Attempt) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
