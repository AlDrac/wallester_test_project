init: copy-env run-prod

run: docker-down docker-up

run-dev: docker-down docker-up-dev

run-prod: docker-down docker-up-prod

copy-env:
	cp .env.dist .env

docker-show:
	docker ps -a

docker-up:
	docker-compose up --build -d

docker-up-dev:
	docker-compose -f docker-compose.yml -f docker-compose-dev.yml up --build -d

docker-up-prod:
	docker-compose -f docker-compose.yml -f docker-compose-prod.yml up --build -d

docker-down:
	docker-compose down -v --remove-orphans

docker-bash-api:
	docker-compose exec -u 0:0 api bash

docker-bash-web:
	docker-compose exec -u 0:0 web bash

docker-bash-migration:
	docker-compose exec -u 1000:1000 migration bash

.DEFAULT_GOAL := init
.PHONY: init run run-dev copy-env docker-up docker-up-dev docker-up-prod docker-down docker-bash-api docker-bash-web docker-bash-migration