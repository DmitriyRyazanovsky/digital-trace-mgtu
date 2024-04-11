package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * UserAnswerAnswer

func userAnswerScan(rows *sql.Rows) (UserAnswer, error) {
	fields := userAnswerStatic()

	err := rows.Scan(
		fields.Id,
		fields.AttemptId,
		fields.QuestionId,
		fields.Answer,
	)
	if err != nil {
		return UserAnswer{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddUserAnswerOut struct {
	UserAnswer UserAnswer
}

func (db *Database) AddUserAnswer(tx *sql.Tx, params UserAnswer) (*AddUserAnswerOut, error) {
	query, args := db.gen.Add(DigitalTraceUserAnswer, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddUserAnswerOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := userAnswerScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.UserAnswer = row

	return out, nil
}

// * FIND
type FindUserAnswerOut struct {
	IsFound    bool
	UserAnswer []UserAnswer
}

func (db *Database) FindUserAnswer(tx *sql.Tx, where UserAnswer) (*FindUserAnswerOut, error) {
	query, args := db.gen.Find(DigitalTraceUserAnswer, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindUserAnswerOut{}

	for rows.Next() {
		row, err := userAnswerScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addUserAnswerScan(rows)]")
		}

		out.UserAnswer = append(out.UserAnswer, row)
	}

	if len(out.UserAnswer) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeUserAnswerOut struct {
	IsFound    bool
	UserAnswer []UserAnswer
}

func (db *Database) ChangeUserAnswer(tx *sql.Tx, set UserAnswer, where UserAnswer) (*ChangeUserAnswerOut, error) {
	query, args := db.gen.Change(DigitalTraceUserAnswer, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeUserAnswerOut{}

	for rows.Next() {
		row, err := userAnswerScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addUserAnswerScan(rows)]")
		}

		out.UserAnswer = append(out.UserAnswer, row)
	}

	if len(out.UserAnswer) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteUserAnswerOut struct {
	IsFound    bool
	UserAnswer []UserAnswer
}

func (db *Database) DeleteUserAnswer(tx *sql.Tx, where UserAnswer) (*DeleteUserAnswerOut, error) {
	query, args := db.gen.Delete(DigitalTraceUserAnswer, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteUserAnswerOut{}

	for rows.Next() {
		row, err := userAnswerScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addUserAnswerScan(rows)]")
		}

		out.UserAnswer = append(out.UserAnswer, row)
	}

	if len(out.UserAnswer) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
