up:
	docker-compose up

upd:
	docker-compose up -d

down:
	docker-compose down

clean:
	@make down
	docker-compose rm

add-hosts:
	cat ./hosts >> /etc/hosts

test:
	go test -race ./...
