package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * Question

func questionScan(rows *sql.Rows) (Question, error) {
	fields := questionStatic()

	err := rows.Scan(
		fields.Id,
		fields.TestId,
		fields.ButtonTypeId,
		fields.Number,
		fields.Content,
		fields.Answer,
	)
	if err != nil {
		return Question{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddQuestionOut struct {
	Question Question
}

func (db *Database) AddQuestion(tx *sql.Tx, params Question) (*AddQuestionOut, error) {
	query, args := db.gen.Add(DigitalTraceQuestion, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddQuestionOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := questionScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.Question = row

	return out, nil
}

// * FIND
type FindQuestionOut struct {
	IsFound  bool
	Question []Question
}

func (db *Database) FindQuestion(tx *sql.Tx, where Question) (*FindQuestionOut, error) {
	query, args := db.gen.Find(DigitalTraceQuestion, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindQuestionOut{}

	for rows.Next() {
		row, err := questionScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addQuestionScan(rows)]")
		}

		out.Question = append(out.Question, row)
	}

	if len(out.Question) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeQuestionOut struct {
	IsFound  bool
	Question []Question
}

func (db *Database) ChangeQuestion(tx *sql.Tx, set Question, where Question) (*ChangeQuestionOut, error) {
	query, args := db.gen.Change(DigitalTraceQuestion, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeQuestionOut{}

	for rows.Next() {
		row, err := questionScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addQuestionScan(rows)]")
		}

		out.Question = append(out.Question, row)
	}

	if len(out.Question) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteQuestionOut struct {
	IsFound  bool
	Question []Question
}

func (db *Database) DeleteQuestion(tx *sql.Tx, where Question) (*DeleteQuestionOut, error) {
	query, args := db.gen.Delete(DigitalTraceQuestion, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteQuestionOut{}

	for rows.Next() {
		row, err := questionScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addQuestionScan(rows)]")
		}

		out.Question = append(out.Question, row)
	}

	if len(out.Question) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
