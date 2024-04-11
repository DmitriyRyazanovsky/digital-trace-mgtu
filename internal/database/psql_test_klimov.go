package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * TestKlimov

func testKlimovScan(rows *sql.Rows) (TestKlimov, error) {
	fields := testKlimovStatic()

	err := rows.Scan(
		fields.Id,
		fields.UserId,
		fields.AttemptId,
		fields.HumanSign,
		fields.HumanHuman,
		fields.HumanNature,
		fields.HumanTechnic,
		fields.HumanSignSystem,
	)
	if err != nil {
		return TestKlimov{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddTestKlimovOut struct {
	TestKlimov TestKlimov
}

func (db *Database) AddTestKlimov(tx *sql.Tx, params TestKlimov) (*AddTestKlimovOut, error) {
	query, args := db.gen.Add(DigitalTraceTestKlimov, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddTestKlimovOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := testKlimovScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.TestKlimov = row

	return out, nil
}

// * FIND
type FindTestKlimovOut struct {
	IsFound    bool
	TestKlimov []TestKlimov
}

func (db *Database) FindTestKlimov(tx *sql.Tx, where TestKlimov) (*FindTestKlimovOut, error) {
	query, args := db.gen.Find(DigitalTraceTestKlimov, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindTestKlimovOut{}

	for rows.Next() {
		row, err := testKlimovScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[testKlimovScan(rows)]")
		}

		out.TestKlimov = append(out.TestKlimov, row)
	}

	if len(out.TestKlimov) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeTestKlimovOut struct {
	IsFound    bool
	TestKlimov []TestKlimov
}

func (db *Database) ChangeTestKlimov(tx *sql.Tx, set TestKlimov, where TestKlimov) (*ChangeTestKlimovOut, error) {
	query, args := db.gen.Change(DigitalTraceTestKlimov, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeTestKlimovOut{}

	for rows.Next() {
		row, err := testKlimovScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[testKlimovScan(rows)]")
		}

		out.TestKlimov = append(out.TestKlimov, row)
	}

	if len(out.TestKlimov) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteTestKlimovOut struct {
	IsFound    bool
	TestKlimov []TestKlimov
}

func (db *Database) DeleteTestKlimov(tx *sql.Tx, where TestKlimov) (*DeleteTestKlimovOut, error) {
	query, args := db.gen.Delete(DigitalTraceTestKlimov, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteTestKlimovOut{}

	for rows.Next() {
		row, err := testKlimovScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[testKlimovScan(rows)]")
		}

		out.TestKlimov = append(out.TestKlimov, row)
	}

	if len(out.TestKlimov) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
