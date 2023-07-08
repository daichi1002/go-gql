gqlgen:
	go get github.com/99designs/gqlgen@latest; \
	go run github.com/99designs/gqlgen generate;

run:
	docker-compose down; \
	docker-compose up -d;

migrate-up:
	migrate --path migrations --database 'mysql://root@tcp(127.0.0.1:3310)/gql?charset=utf8&parseTime=True&loc=Local&timeout=10s' -verbose up;

mockgen:
	mockgen -source usecases/usecases_interfaces.go -destination usecases/mock/mock_usecases.go -package mock_usecases;
	mockgen -source adapters/repositories/repositories_interfaces.go -destination adapters/repositories/mock/mock_repositories.go -package mock_repositories;
	mockgen -source adapters/interfaces.go -destination infra/mock/mock_infra.go -package mock_infra;

test:
	go test -v ./...;