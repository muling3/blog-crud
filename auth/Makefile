.PHONY: gen migrations-up migrations-down start-db stop-db runserver

gen:
	sqlc generate
migrations-up:
	migrate -path "db/schema" -database "postgresql://root:password@localhost:5432/blog_microservice?sslmode=disable" up
migrations-down:
	migrate -path "db/schema" -database "postgresql://root:password@localhost:5432/blog_microservice?sslmode=disable" down
start-db:
	docker start database-post
stop-db:
	docker stop database-post
runserver:
	go run main.go