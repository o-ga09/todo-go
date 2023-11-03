DOCKER_TAG := latest

run:
	DBUSER=todo DBPASSWORD=password DBHOST=localhost DBPORT=3306 DBNAME=todo ENV=dev go run ./cmd/.
up:
	docker compose up -d
down:
	docker compose down
build:
	docker build --platform linux/amd64 -t taiti09/mah-api:${DOCKER_TAG} --target deploy ./
build-local:
	docker compose build --no-cache
logs:
	docker compose logs -f
rm:
	docker compose down all --volumes --remove-orphans
test:
	go test -race ./...



# curl -X POST -H 'Content-type: application/json' -d '{"name":"test","password":"dmwmdodk"}' http://localhost:8080/api/v1/user