package database

import (
	"database/sql"

	"github.com/pkg/errors"
)

func roleScan(rows *sql.Rows) (Role, error) {
	role := roleStatic()

	err := rows.Scan(
		role.Id,
		role.Name,
	)
	if err != nil {
		return Role{}, errors.Wrapf(err, "query error: [rows.Scan()]")
	}

	return role, nil
}

// * ADD
type AddRoleOut struct {
	Role Role
}

func (db *Database) AddRole(tx *sql.Tx, role Role) (*AddRoleOut, error) {
	query, args := db.gen.Add(DigitalTraceRole, role)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &AddRoleOut{}

	if !rows.Next() {
		return nil, errors.New("unable to add params")
	}

	row, err := roleScan(rows)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [rows.Scan(db.gen.Scan())]")
	}

	out.Role = row

	return out, nil
}

// * FIND
type FindRoleOut struct {
	IsFound bool
	Role    []Role
}

func (db *Database) FindRole(tx *sql.Tx, where Role) (*FindRoleOut, error) {
	query, args := db.gen.Find(DigitalTraceRole, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &FindRoleOut{}

	for rows.Next() {
		row, err := roleScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addRoleScan(rows)]")
		}

		out.Role = append(out.Role, row)
	}

	if len(out.Role) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * CHANGE
type ChangeRoleOut struct {
	IsFound bool
	Role    []Role
}

func (db *Database) ChangeRole(tx *sql.Tx, set Role, where Role) (*ChangeRoleOut, error) {
	query, args := db.gen.Change(DigitalTraceRole, set, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &ChangeRoleOut{}

	for rows.Next() {
		row, err := roleScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addRoleScan(rows)]")
		}

		out.Role = append(out.Role, row)
	}

	if len(out.Role) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}

// * DELETE
type DeleteRoleOut struct {
	IsFound bool
	Role    []Role
}

func (db *Database) DeleteRole(tx *sql.Tx, where Role) (*DeleteRoleOut, error) {
	query, args := db.gen.Delete(DigitalTraceRole, where)

	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "query error: [tx.Query]")
	}
	defer rows.Close()

	out := &DeleteRoleOut{}

	for rows.Next() {
		row, err := roleScan(rows)
		if err != nil {
			return nil, errors.Wrapf(err, "[addRoleScan(rows)]")
		}

		out.Role = append(out.Role, row)
	}

	if len(out.Role) == 0 {
		return out, nil
	}

	out.IsFound = true

	return out, nil
}
