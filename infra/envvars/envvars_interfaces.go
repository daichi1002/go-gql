package envvars

type EnvironmentVariablesInterface interface {
	GetMysqlUser() string
	GetMysqlHost() string
	GetMysqlPort() string
	GetMysqlPassword() string
	GetMysqlDB() string
}
