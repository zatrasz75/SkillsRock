run:
	go run cmd/app/main.go

up:
	sql-migrate new up

down:
	sql-migrate down

swag:
	swag init -d internal/handlers/ -g router.go --parseDependency --parseDepth 3

dock:
	docker compose up -d