build:
	docker-compose -f ./deploy/docker-compose.yml build --no-cache generator

run:
	docker-compose -f ./deploy/docker-compose.yml up --remove-orphans generator

test:
	docker-compose -f ./deploy/docker-compose.yml up tests

lint:
	docker-compose -f ./deploy/docker-compose.yml up linter

down:
	docker-compose -f deploy/docker-compose.yml down