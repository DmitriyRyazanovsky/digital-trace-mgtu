package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * ButtonType

func buttonTypeScan(rows *sql.Rows) (ButtonType, error) {
	fields := buttonTypeStatic()

	err := rows.Scan(
		fields.Id,
		fields.Name,
	)
	if err != nil {
		return ButtonType{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddButtonTypeOut struct {
	ButtonType ButtonType
}

func (db *Database) AddButtonType(tx *sql.Tx, params ButtonType) (*AddButtonTypeOut, error) {
	query, args := db.gen.Add(DigitalTraceButtonType, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddButtonTypeOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := buttonTypeScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.ButtonType = row

	return out, nil
}

// * FIND
type FindButtonTypeOut struct {
	IsFound    bool
	ButtonType []ButtonType
}

func (db *Database) FindButtonType(tx *sql.Tx, where ButtonType) (*FindButtonTypeOut, error) {
	query, args := db.gen.Find(DigitalTraceButtonType, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindButtonTypeOut{}

	for rows.Next() {
		row, err := buttonTypeScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addButtonTypeScan(rows)]")
		}

		out.ButtonType = append(out.ButtonType, row)
	}

	if len(out.ButtonType) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeButtonTypeOut struct {
	IsFound    bool
	ButtonType []ButtonType
}

func (db *Database) ChangeButtonType(tx *sql.Tx, set ButtonType, where ButtonType) (*ChangeButtonTypeOut, error) {
	query, args := db.gen.Change(DigitalTraceButtonType, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeButtonTypeOut{}

	for rows.Next() {
		row, err := buttonTypeScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addButtonTypeScan(rows)]")
		}

		out.ButtonType = append(out.ButtonType, row)
	}

	if len(out.ButtonType) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteButtonTypeOut struct {
	IsFound    bool
	ButtonType []ButtonType
}

func (db *Database) DeleteButtonType(tx *sql.Tx, where ButtonType) (*DeleteButtonTypeOut, error) {
	query, args := db.gen.Delete(DigitalTraceButtonType, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteButtonTypeOut{}

	for rows.Next() {
		row, err := buttonTypeScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addButtonTypeScan(rows)]")
		}

		out.ButtonType = append(out.ButtonType, row)
	}

	if len(out.ButtonType) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
