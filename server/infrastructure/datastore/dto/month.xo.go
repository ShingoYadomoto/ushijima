// Package dto contains the types for schema 'public'.
package dto

// Code generated by xo. DO NOT EDIT.

import (
	"gopkg.in/guregu/null.v3"
)

// Month represents a row from 'public.months'.
type Month struct {
	ID         int         `db:"id"`          // id
	Display    null.String `db:"display"`     // display
	CreateDate null.Time   `db:"create_date"` // create_date
	UpdateDate null.Time   `db:"update_date"` // update_date
}
