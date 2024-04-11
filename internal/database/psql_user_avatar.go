package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * UserAvatarAnswer

func userAvatarScan(rows *sql.Rows) (UserAvatar, error) {
	fields := userAvatarStatic()

	err := rows.Scan(
		fields.Id,
		fields.UserId,
		fields.Prefix,
		fields.Path,
	)
	if err != nil {
		return UserAvatar{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddUserAvatarOut struct {
	UserAvatar UserAvatar
}

func (db *Database) AddUserAvatar(tx *sql.Tx, params UserAvatar) (*AddUserAvatarOut, error) {
	query, args := db.gen.Add(DigitalTraceUserAvatar, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddUserAvatarOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := userAvatarScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.UserAvatar = row

	return out, nil
}

// * FIND
type FindUserAvatarOut struct {
	IsFound    bool
	UserAvatar []UserAvatar
}

func (db *Database) FindUserAvatar(tx *sql.Tx, where UserAvatar) (*FindUserAvatarOut, error) {
	query, args := db.gen.Find(DigitalTraceUserAvatar, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindUserAvatarOut{}

	for rows.Next() {
		row, err := userAvatarScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[userAvatarScan(rows)]")
		}

		out.UserAvatar = append(out.UserAvatar, row)
	}

	if len(out.UserAvatar) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeUserAvatarOut struct {
	IsFound    bool
	UserAvatar []UserAvatar
}

func (db *Database) ChangeUserAvatar(tx *sql.Tx, set UserAvatar, where UserAvatar) (*ChangeUserAvatarOut, error) {
	query, args := db.gen.Change(DigitalTraceUserAvatar, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeUserAvatarOut{}

	for rows.Next() {
		row, err := userAvatarScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[userAvatarScan(rows)]")
		}

		out.UserAvatar = append(out.UserAvatar, row)
	}

	if len(out.UserAvatar) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteUserAvatarOut struct {
	IsFound    bool
	UserAvatar []UserAvatar
}

func (db *Database) DeleteUserAvatar(tx *sql.Tx, where UserAvatar) (*DeleteUserAvatarOut, error) {
	query, args := db.gen.Delete(DigitalTraceUserAvatar, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteUserAvatarOut{}

	for rows.Next() {
		row, err := userAvatarScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[userAvatarScan(rows)]")
		}

		out.UserAvatar = append(out.UserAvatar, row)
	}

	if len(out.UserAvatar) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
