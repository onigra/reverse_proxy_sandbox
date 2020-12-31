up:
	docker-compose up

down:
	docker-compose down

clean:
	@make down
	docker-compose rm

add-hosts:
	cat ./hosts >> /etc/hosts
