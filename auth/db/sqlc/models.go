// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Type string

const (
	TypeAdmin   Type = "admin"
	TypeBlogger Type = "blogger"
	TypeRegular Type = "regular"
)

func (e *Type) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Type(s)
	case string:
		*e = Type(s)
	default:
		return fmt.Errorf("unsupported scan type for Type: %T", src)
	}
	return nil
}

type NullType struct {
	Type  Type
	Valid bool // Valid is true if Type is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullType) Scan(value interface{}) error {
	if value == nil {
		ns.Type, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Type.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Type), nil
}

type User struct {
	ID int32 `json:"id"`
	// User first + last names
	FullName  string        `json:"full_name"`
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Active    sql.NullInt32 `json:"active"`
	Type      NullType      `json:"type"`
	LastLogin sql.NullTime  `json:"last_login"`
}