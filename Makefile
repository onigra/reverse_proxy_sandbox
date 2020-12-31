up:
	docker-compose up

down:
	docker-compose down

clean:
	@make down
	docker-compose rm
