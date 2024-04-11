package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

func testScan(rows *sql.Rows) (Test, error) {
	fields := testStatic()

	err := rows.Scan(
		fields.Id,
		fields.Name,
		fields.PgName,
		fields.Content,
		fields.Description,
	)
	if err != nil {
		return Test{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddTestOut struct {
	Test Test
}

func (db *Database) AddTest(tx *sql.Tx, params Test) (*AddTestOut, error) {
	query, args := db.gen.Add(DigitalTraceTest, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddTestOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := testScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.Test = row

	return out, nil
}

// * FIND
type FindTestOut struct {
	IsFound bool
	Test    []Test
}

func (db *Database) FindTest(tx *sql.Tx, where Test) (*FindTestOut, error) {
	query, args := db.gen.Find(DigitalTraceTest, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindTestOut{}

	for rows.Next() {
		row, err := testScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[findTestScan(rows)]")
		}

		out.Test = append(out.Test, row)
	}

	if len(out.Test) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeTestOut struct {
	IsFound bool
	Test    []Test
}

func (db *Database) ChangeTest(tx *sql.Tx, set Test, where Test) (*ChangeTestOut, error) {
	query, args := db.gen.Change(DigitalTraceTest, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeTestOut{}

	for rows.Next() {
		row, err := testScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[changeTestScan(rows)]")
		}

		out.Test = append(out.Test, row)
	}

	if len(out.Test) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteTestOut struct {
	IsFound bool
	Test    []Test
}

func (db *Database) DeleteTest(tx *sql.Tx, where Test) (*DeleteTestOut, error) {
	query, args := db.gen.Delete(DigitalTraceTest, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteTestOut{}

	for rows.Next() {
		row, err := testScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[deleteTestScan(rows)]")
		}

		out.Test = append(out.Test, row)
	}

	if len(out.Test) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
