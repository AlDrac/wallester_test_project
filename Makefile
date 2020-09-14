init: copy-env run

run: docker-down docker-up

copy-env:
	cp .env.dist .env

docker-up:
	docker-compose up --build -d

docker-down:
	docker-compose down -v --remove-orphans

docker-bash-api:
	docker-compose exec -u 0:0 api bash

docker-bash-web:
	docker-compose exec -u 0:0 web bash

docker-bash-migration:
	docker-compose exec -u 0:0 migration bash

.DEFAULT_GOAL := init
.PHONY:init run copy-env docker-up docker-down docker-bash-api docker-bash-web docker-bash-migration