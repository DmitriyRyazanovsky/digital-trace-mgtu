package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

// * User

func userScan(rows *sql.Rows) (User, error) {
	fields := userStatic()

	err := rows.Scan(
		fields.Id,
		fields.CreatedAt,
		fields.UpdatedAt,
		fields.RoleId,
		fields.Email,
		fields.Login,
		fields.Name,
		fields.Surname,
		fields.Password,
	)
	if err != nil {
		return User{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return fields, nil
}

// * ADD
type AddUserOut struct {
	User User
}

func (db *Database) AddUser(tx *sql.Tx, params User) (*AddUserOut, error) {
	query, args := db.gen.Add(DigitalTraceUser, params)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddUserOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := userScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.User = row

	return out, nil
}

// * FIND
type FindUserOut struct {
	IsFound bool
	User    []User
}

func (db *Database) FindUser(tx *sql.Tx, where User) (*FindUserOut, error) {
	query, args := db.gen.Find(DigitalTraceUser, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindUserOut{}

	for rows.Next() {
		row, err := userScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[userScan(rows)]")
		}

		out.User = append(out.User, row)
	}

	if len(out.User) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeUserOut struct {
	IsFound bool
	User    []User
}

func (db *Database) ChangeUser(tx *sql.Tx, set User, where User) (*ChangeUserOut, error) {
	query, args := db.gen.Change(DigitalTraceUser, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeUserOut{}

	for rows.Next() {
		row, err := userScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addUserScan(rows)]")
		}

		out.User = append(out.User, row)
	}

	if len(out.User) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteUserOut struct {
	IsFound bool
	User    []User
}

func (db *Database) DeleteUser(tx *sql.Tx, where User) (*DeleteUserOut, error) {
	query, args := db.gen.Delete(DigitalTraceUser, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteUserOut{}

	for rows.Next() {
		row, err := userScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addUserScan(rows)]")
		}

		out.User = append(out.User, row)
	}

	if len(out.User) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
