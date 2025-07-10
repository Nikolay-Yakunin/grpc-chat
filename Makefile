CC=go




all: 

auth_build:
	docker-compose -f docker-compose.yaml build auth-service

auth_run:
	docker-compose -f docker-compose.yaml up auth-service

auth_stop:
	docker-compose -f docker-compose.yaml down

auth_gen_migrate:
	docker run --rm -v $(shell pwd)/database/migrations:/migrations migrate/migrate \
    create -ext sql -dir /migrations -seq create_auth_table

auth_migrate:
	docker run --rm -v $(shell pwd)/database/migrations:/migrations --network host migrate/migrate \
    -path=/migrations -database "postgres://user:password@localhost:5432/authdb?sslmode=disable" up
