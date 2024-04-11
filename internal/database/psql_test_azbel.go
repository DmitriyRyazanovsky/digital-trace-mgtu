package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * TestAzbel

func testAzbelScan(rows *sql.Rows) (TestAzbel, error) {
	fields := testAzbelStatic()

	err := rows.Scan(
		fields.Id,
		fields.UserId,
		fields.AttemptId,
		fields.PhysicsMaths,
		fields.BiologyChemistry,
		fields.Tourism,
		fields.Medicine,
		fields.InformationTechnology,
		fields.Construction,
		fields.EngineeringAndTechnicalSphere,
		fields.EconomicsFinance,
		fields.BusinessManagement,
		fields.ForeignLanguages,
		fields.TransportLogistics,
		fields.StrongStructure,
		fields.SocioPoliticalSphere,
		fields.Journalism,
		fields.Jurisprudence,
		fields.Education,
		fields.ServiceSectorTrade,
		fields.PhysicalEducationAndSports,
		fields.MusicalAndPerformingArts,
		fields.FineArtDesign,
	)
	if err != nil {
		return TestAzbel{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddTestAzbelOut struct {
	TestAzbel TestAzbel
}

func (db *Database) AddTestAzbel(tx *sql.Tx, params TestAzbel) (*AddTestAzbelOut, error) {
	query, args := db.gen.Add(DigitalTraceTestAzbel, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddTestAzbelOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := testAzbelScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.TestAzbel = row

	return out, nil
}

// * FIND
type FindTestAzbelOut struct {
	IsFound   bool
	TestAzbel []TestAzbel
}

func (db *Database) FindTestAzbel(tx *sql.Tx, where TestAzbel) (*FindTestAzbelOut, error) {
	query, args := db.gen.Find(DigitalTraceTestAzbel, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindTestAzbelOut{}

	for rows.Next() {
		row, err := testAzbelScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addTestAzbelScan(rows)]")
		}

		out.TestAzbel = append(out.TestAzbel, row)
	}

	if len(out.TestAzbel) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeTestAzbelOut struct {
	IsFound   bool
	TestAzbel []TestAzbel
}

func (db *Database) ChangeTestAzbel(tx *sql.Tx, set TestAzbel, where TestAzbel) (*ChangeTestAzbelOut, error) {
	query, args := db.gen.Change(DigitalTraceTestAzbel, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeTestAzbelOut{}

	for rows.Next() {
		row, err := testAzbelScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addTestAzbelScan(rows)]")
		}

		out.TestAzbel = append(out.TestAzbel, row)
	}

	if len(out.TestAzbel) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteTestAzbelOut struct {
	IsFound   bool
	TestAzbel []TestAzbel
}

func (db *Database) DeleteTestAzbel(tx *sql.Tx, where TestAzbel) (*DeleteTestAzbelOut, error) {
	query, args := db.gen.Delete(DigitalTraceTestAzbel, where)
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteTestAzbelOut{}

	for rows.Next() {
		row, err := testAzbelScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addTestAzbelScan(rows)]")
		}

		out.TestAzbel = append(out.TestAzbel, row)
	}

	if len(out.TestAzbel) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
