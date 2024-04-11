package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * Achievement

func achievementScan(rows *sql.Rows) (Achievement, error) {
	fields := achievementStatic()
	err := rows.Scan(
		fields.Id,
		fields.UserId,
		fields.AchievementID,
		fields.AchievementTypes,
	)
	if err != nil {
		return Achievement{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddAchievementOut struct {
	Achievement Achievement
}

func (db *Database) AddAchievement(tx *sql.Tx, params Achievement) (*AddAchievementOut, error) {
	query, args := db.gen.Add(DigitalTraceAchievement, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddAchievementOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := achievementScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "achievementScan(rows)]")
	}

	out.Achievement = row

	return out, nil
}

// * FIND
type FindAchievementOut struct {
	IsFound     bool
	Achievement []Achievement
}

func (db *Database) FindAchievement(tx *sql.Tx, where Achievement) (*FindAchievementOut, error) {
	query, args := db.gen.Find(DigitalTraceAchievement, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindAchievementOut{}

	for rows.Next() {
		row, err := achievementScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[achievementScan(rows)]")
		}

		out.Achievement = append(out.Achievement, row)
	}

	if len(out.Achievement) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeAchievementOut struct {
	IsFound     bool
	Achievement []Achievement
}

func (db *Database) ChangeAchievement(tx *sql.Tx, set Achievement, where Achievement) (*ChangeAchievementOut, error) {
	query, args := db.gen.Change(DigitalTraceAchievement, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeAchievementOut{}

	for rows.Next() {
		row, err := achievementScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[achievementScan(rows)]")
		}

		out.Achievement = append(out.Achievement, row)
	}

	if len(out.Achievement) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteAchievementOut struct {
	IsFound     bool
	Achievement []Achievement
}

func (db *Database) DeleteAchievement(tx *sql.Tx, where Achievement) (*DeleteAchievementOut, error) {
	query, args := db.gen.Delete(DigitalTraceAchievement, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteAchievementOut{}

	for rows.Next() {
		row, err := achievementScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[achievementScan(rows)]")
		}

		out.Achievement = append(out.Achievement, row)
	}

	if len(out.Achievement) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
