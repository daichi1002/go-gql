gqlgen:
	go mod tidy; \
	go run github.com/99designs/gqlgen generate;

start:
	docker-compose down; \
	docker-compose up -d;