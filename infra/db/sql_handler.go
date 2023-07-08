package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/daichi1002/go-graphql/adapters"
	"github.com/daichi1002/go-graphql/constants"
	"github.com/daichi1002/go-graphql/log"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlHandlerParamsGetter interface {
	GetMysqlUser() string
	GetMysqlHost() string
	GetMysqlPort() string
	GetMysqlPassword() string
	GetMysqlDB() string
}

func NewSqlHandler(params SqlHandlerParamsGetter) adapters.SqlHandler {
	conn, err := sql.Open("mysql", getConnectionString(
		params.GetMysqlUser(), params.GetMysqlHost(), params.GetMysqlPort(), params.GetMysqlPassword(), params.GetMysqlDB(),
	))

	if err != nil {
		log.Fatal(err.Error())
	}

	return &SqlHandler{conn}
}

func getConnectionString(user string, host string, port string, password string, db string) string {
	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=10s", user, host, port, db)

	return dsn
}

func NewSqlHandlerOfDB(db *sql.DB) *SqlHandler {
	return &SqlHandler{db}
}

type SqlResult struct {
	Result sql.Result
}

func (handler *SqlHandler) Execute(ctx context.Context, statement string, args ...interface{}) (adapters.Result, error) {
	tx, ok := ctx.Value(constants.TxCtxKey).(*sql.Tx)

	if ok {
		result, err := tx.Exec(statement, args...)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		return SqlResult{result}, nil
	}

	result, err := handler.Conn.Exec(statement, args...)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return SqlResult{result}, nil
}

func (r SqlResult) LastInsertId() (int64, error) {
	lastInsertId, err := r.Result.LastInsertId()

	if err != nil {
		log.Error(err.Error())
	}

	return lastInsertId, err
}

func (r SqlResult) RowsAffected() (int64, error) {
	rowsAffected, err := r.Result.RowsAffected()

	if err != nil {
		log.Error(err.Error())
	}

	return rowsAffected, err
}

type SqlRow struct {
	Rows *sql.Rows
}

func (handler *SqlHandler) Query(ctx context.Context, statement string, args ...interface{}) (adapters.Rows, error) {
	tx, ok := ctx.Value(constants.TxCtxKey).(*sql.Tx)

	if ok {
		rows, err := tx.Query(statement, args...)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		return SqlRow{rows}, nil
	}
	rows, err := handler.Conn.Query(statement, args...)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return SqlRow{rows}, nil
}

func (r SqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)

	if err != nil {
		log.Error(err.Error())
	}

	return err
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	err := r.Rows.Close()

	if err != nil {
		log.Error(err.Error())
	}

	return err
}

func (handler *SqlHandler) Begin() (*sql.Tx, error) {
	return handler.Conn.Begin()
}
