package envvars

import "os"

type Envvars struct {
	MysqlUser     string
	MysqlHost     string
	MysqlPort     string
	MysqlPassword string
	MysqlDB       string
}

var e = &Envvars{}

func NewEnvVar() {
	e.MysqlUser = os.Getenv("MYSQL_USER")
	e.MysqlHost = os.Getenv("MYSQL_HOST")
	e.MysqlPort = os.Getenv("MYSQL_PORT")
	e.MysqlPassword = os.Getenv("MYSQL_PASSWORD")
	e.MysqlDB = os.Getenv("MYSQL_DB")
}

func (e *Envvars) GetMysqlUser() string {
	return e.MysqlUser
}

func (e *Envvars) GetMysqlHost() string {
	return e.MysqlHost
}

func (e *Envvars) GetMysqlPort() string {
	return e.MysqlPort
}

func (e *Envvars) GetMysqlPassword() string {
	return e.MysqlPassword
}

func (e *Envvars) GetMysqlDB() string {
	return e.MysqlDB
}
