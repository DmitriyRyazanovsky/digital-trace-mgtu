package database

import (
	"fmt"
	"reflect"
)

type Generator struct {
}

// * CHANGE
const (
	changeQuery = "UPDATE %s %s %s RETURNING *;"
)

func (c *Generator) Change(tableName string, set interface{}, where interface{}) (string, []any) {
	args := []any{}

	setQuery := ""

	setN := 0

	paramsType := reflect.TypeOf(set)
	paramsValue := reflect.ValueOf(set)
	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		value := paramsValue.Field(i)

		if value.IsNil() {
			continue
		}

		tag := field.Tag.Get("psql")
		if tag == "" {
			continue
		}

		if setN != 0 {
			setQuery += ", "
		} else {
			setQuery += "SET "
		}

		setQuery += fmt.Sprintf("%s = $%d", tag, setN+1)
		args = append(args, value.Interface())
		setN++
	}

	whereQuery := ""

	whereN := 0

	paramsType = reflect.TypeOf(where)
	paramsValue = reflect.ValueOf(where)
	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		value := paramsValue.Field(i)

		if value.IsNil() {
			continue
		}

		tag := field.Tag.Get("psql")
		if tag == "" {
			continue
		}

		if whereN != 0 {
			whereQuery += " AND "
		} else {
			whereQuery += "WHERE "
		}

		whereQuery += fmt.Sprintf("%s = $%d", tag, setN+whereN+1)
		args = append(args, value.Interface())
		whereN++
	}

	return fmt.Sprintf(changeQuery, tableName, setQuery, whereQuery), args
}

// * FIND
const (
	findQuery = "SELECT * FROM %s %s;"
)

func (c *Generator) Find(tableName string, where interface{}) (string, []any) {
	args := []any{}

	whereQuery := ""

	whereN := 0

	paramsType := reflect.TypeOf(where)
	paramsValue := reflect.ValueOf(where)
	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		value := paramsValue.Field(i)

		if value.IsNil() {
			continue
		}

		tag := field.Tag.Get("psql")
		if tag == "" {
			continue
		}

		if whereN != 0 {
			whereQuery += " AND "
		} else {
			whereQuery += "WHERE "
		}

		whereQuery += fmt.Sprintf("%s = $%d", tag, whereN+1)
		args = append(args, value.Interface())
		whereN++
	}

	return fmt.Sprintf(findQuery, tableName, whereQuery), args
}

// * ADD
const (
	addQuery = "INSERT INTO %s (%s) VALUES (%s) RETURNING *"
)

func (c *Generator) Add(tableName string, params interface{}) (string, []any) {
	args := []any{}

	intoQuery := ""
	valuesQuery := ""

	n := 0

	paramsType := reflect.TypeOf(params)
	paramsValue := reflect.ValueOf(params)
	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		value := paramsValue.Field(i)

		if value.IsNil() {
			continue
		}

		tag := field.Tag.Get("psql")
		if tag == "" {
			continue
		}

		if n != 0 {
			intoQuery += ", "
			valuesQuery += ", "
		}

		intoQuery += tag
		valuesQuery += fmt.Sprintf("$%d", n+1)

		args = append(args, value.Interface())
		n++
	}

	return fmt.Sprintf(addQuery, tableName, intoQuery, valuesQuery), args
}

// * DELETE
const (
	deleteQuery = "DELETE FROM %s WHERE %s RETURNING *"
)

func (c *Generator) Delete(tableName string, where interface{}) (string, []any) {
	args := []any{}

	whereQuery := ""

	n := 0

	paramsType := reflect.TypeOf(where)
	paramsValue := reflect.ValueOf(where)
	for i := 0; i < paramsType.NumField(); i++ {
		field := paramsType.Field(i)
		value := paramsValue.Field(i)

		if value.IsNil() {
			continue
		}

		tag := field.Tag.Get("psql")
		if tag == "" {
			continue
		}

		if n != 0 {
			whereQuery += " AND "
		}

		args = append(args, value.Interface())
		whereQuery += fmt.Sprintf("%s = $%d", tag, n+1)
		n++
	}

	return fmt.Sprintf(deleteQuery, tableName, whereQuery), args
}
