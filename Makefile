start:
	docker-compose up -d

start-build:
	docker-compose up -d --build

stop:
	docker-compose down

restart:
	make stop && make start

restart-build:
	make stop && make start-build

# dev:
# 	air
dev:
	APP_SERVICE=api go run main.go

run:
	go run main.go

install:
	export GOPRIVATE=gitlab.finema.co/finema/* && git config --global url."git@gitlab.finema.co:".insteadOf "https://gitlab.finema.co/" && go get

logs:
	 docker logs -f api

# migrate:
# 	docker-compose up -d --build migration
migrate:
	APP_SERVICE=migration go run main.go

test:
	go test ./...
