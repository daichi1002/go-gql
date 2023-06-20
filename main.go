package main

import (
	"github.com/daichi1002/go-graphql/di"
	"github.com/daichi1002/go-graphql/infra/envvars"
	"github.com/daichi1002/go-graphql/server"
)

func main() {
	// 環境変数の読み込み
	env := envvars.GetInstance()
	di.Do(env)

	// DB接続
	server.Serve()
}
