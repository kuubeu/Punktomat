build-dev:
	docker-compose -f docker-compose.yml build

dev:
	docker-compose -f docker-compose.yml up

build-prod:
	docker-compose -f docker-compose.prod.yml build --no-cache

prod:
	docker-compose -f docker-compose.prod.yml up

