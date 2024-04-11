package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * AchievementType

func achievementTypeScan(rows *sql.Rows) (AchievementType, error) {
	fields := achievementTypeStatic()

	err := rows.Scan(
		fields.Id,
		fields.Name,
	)
	if err != nil {
		return AchievementType{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddAchievementTypeOut struct {
	AchievementType AchievementType
}

func (db *Database) AddAchievementType(tx *sql.Tx, params AchievementType) (*AddAchievementTypeOut, error) {
	query, args := db.gen.Add(DigitalTraceAchievementType, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddAchievementTypeOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := achievementTypeScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.AchievementType = row

	return out, nil
}

// * FIND
type FindAchievementTypeOut struct {
	IsFound         bool
	AchievementType []AchievementType
}

func (db *Database) FindAchievementType(tx *sql.Tx, where AchievementType) (*FindAchievementTypeOut, error) {
	query, args := db.gen.Find(DigitalTraceAchievementType, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindAchievementTypeOut{}

	for rows.Next() {
		row, err := achievementTypeScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addAchievementTypeScan(rows)]")
		}

		out.AchievementType = append(out.AchievementType, row)
	}

	if len(out.AchievementType) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeAchievementTypeOut struct {
	IsFound         bool
	AchievementType []AchievementType
}

func (db *Database) ChangeAchievementType(tx *sql.Tx, set AchievementType, where AchievementType) (*ChangeAchievementTypeOut, error) {
	query, args := db.gen.Change(DigitalTraceAchievementType, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeAchievementTypeOut{}

	for rows.Next() {
		row, err := achievementTypeScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addAchievementTypeScan(rows)]")
		}

		out.AchievementType = append(out.AchievementType, row)
	}

	if len(out.AchievementType) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteAchievementTypeOut struct {
	IsFound         bool
	AchievementType []AchievementType
}

func (db *Database) DeleteAchievementType(tx *sql.Tx, where AchievementType) (*DeleteAchievementTypeOut, error) {
	query, args := db.gen.Delete(DigitalTraceAchievementType, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteAchievementTypeOut{}

	for rows.Next() {
		row, err := achievementTypeScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addAchievementTypeScan(rows)]")
		}

		out.AchievementType = append(out.AchievementType, row)
	}

	if len(out.AchievementType) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
