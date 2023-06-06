include .env
export

test:
	go clean --testcache; go test -v ./...

rundb:
	docker run --name urlsdb -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=urls -p 5432:5432 -d postgres 

migrate-up:
	migrate -source file:migrations -database 'postgres://postgres:postgres@localhost?sslmode=disable' -verbose up

migrate-down:
	migrate -source file:migrations -database 'postgres://postgres:postgres@localhost?sslmode=disable' -verbose down