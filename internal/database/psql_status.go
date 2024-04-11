package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

func statusScan(rows *sql.Rows) (Status, error) {
	fields := statusStatic()

	err := rows.Scan(
		fields.Id,
		fields.Name,
	)
	if err != nil {
		return Status{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddStatusOut struct {
	Status Status
}

func (db *Database) AddStatus(tx *sql.Tx, params Status) (*AddStatusOut, error) {
	query, args := db.gen.Add(DigitalTraceStatus, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddStatusOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := statusScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.Status = row

	return out, nil
}

// * FIND
type FindStatusOut struct {
	IsFound bool
	Status  []Status
}

func (db *Database) FindStatus(tx *sql.Tx, where Status) (*FindStatusOut, error) {
	query, args := db.gen.Find(DigitalTraceStatus, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindStatusOut{}

	for rows.Next() {
		row, err := statusScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addStatusScan(rows)]")
		}

		out.Status = append(out.Status, row)
	}

	if len(out.Status) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeStatusOut struct {
	IsFound bool
	Status  []Status
}

func (db *Database) ChangeStatus(tx *sql.Tx, set Status, where Status) (*ChangeStatusOut, error) {
	query, args := db.gen.Change(DigitalTraceStatus, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeStatusOut{}

	for rows.Next() {
		row, err := statusScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addStatusScan(rows)]")
		}

		out.Status = append(out.Status, row)
	}

	if len(out.Status) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteStatusOut struct {
	IsFound bool
	Status  []Status
}

func (db *Database) DeleteStatus(tx *sql.Tx, where Status) (*DeleteStatusOut, error) {
	query, args := db.gen.Delete(DigitalTraceStatus, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteStatusOut{}

	for rows.Next() {
		row, err := statusScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addStatusScan(rows)]")
		}

		out.Status = append(out.Status, row)
	}

	if len(out.Status) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
