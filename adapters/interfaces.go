package adapters

import (
	"context"
	"database/sql"
)

type SqlHandler interface {
	Execute(context.Context, string, ...interface{}) (Result, error)
	Query(context.Context, string, ...interface{}) (Rows, error)
	Begin() (*sql.Tx, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
