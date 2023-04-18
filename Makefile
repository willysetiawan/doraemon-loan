run_local:
	go run main.go

run_migrate:
	go run ./db/migrate/migrate.go

run_seed:
	psql -U postgres base_code < db/seed.sql

run_docker:
	docker stop basecodeapiserver || true && docker rm basecodeapiserver || true
	docker build --tag basecode-api:dev .
	docker run --name basecodeapiserver -d -p 4000:4000 basecode-api:dev
