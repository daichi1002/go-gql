gqlgen:
	go get github.com/99designs/gqlgen@latest; \
	go run github.com/99designs/gqlgen generate;

run:
	docker-compose down; \
	docker-compose up -d;

migrate-up:
	migrate --path migrations --database 'mysql://root@tcp(127.0.0.1:3310)/gql?charset=utf8&parseTime=True&loc=Local&timeout=10s' -verbose up;