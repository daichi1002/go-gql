gqlgen:
	go mod tidy; \
	go run github.com/99designs/gqlgen generate;

wire:
	cd di; \
	wire