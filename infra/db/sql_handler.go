package db

import (
	"database/sql"
	"fmt"

	"github.com/daichi1002/go-graphql/adapters"
	"github.com/daichi1002/go-graphql/util"
	_ "github.com/go-sql-driver/mysql"
)

var logger = util.NewLogger()

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
		logger.Fatal(err)
	}

	return &SqlHandler{conn}
}

func getConnectionString(user string, host string, port string, password string, db string) string {
	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=10s", user, host, port, db)

	return dsn
}

type SqlResult struct {
	Result sql.Result
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (adapters.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)

	if err != nil {
		return nil, err
	}

	res.Result = result
	return res, nil
}

func (r SqlResult) LastInsertId() (int64, error) {
	lastInsertId, err := r.Result.LastInsertId()

	if err != nil {
		logger.Error(err)
	}

	return lastInsertId, err
}

func (r SqlResult) RowsAffected() (int64, error) {
	rowsAffected, err := r.Result.RowsAffected()

	if err != nil {
		logger.Error(err)
	}

	return rowsAffected, err
}

type SqlRow struct {
	Rows *sql.Rows
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (adapters.Rows, error) {
	rows, err := handler.Conn.Query(statement, args...)

	if err != nil {
		return nil, err
	}

	return SqlRow{rows}, nil
}

func (r SqlRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)

	if err != nil {
		logger.Error(err)
	}

	return err
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	err := r.Rows.Close()

	if err != nil {
		logger.Error(err)
	}

	return err
}
